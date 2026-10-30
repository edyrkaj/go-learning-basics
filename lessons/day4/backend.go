package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
)

// ErrNotFound is returned when a job id is missing from the store.
var ErrNotFound = errors.New("job not found")

// Job is one stored piece of work.
type Job struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Store persists jobs. Put returns the job with the id the store assigned.
type Store interface {
	Put(ctx context.Context, job Job) (Job, error)
	Get(ctx context.Context, id string) (Job, error)
}

// MemoryStore is an in-memory Store guarded by a mutex.
type MemoryStore struct {
	mu   sync.Mutex
	jobs map[string]Job
	next int
}

// NewMemoryStore returns an empty store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{jobs: make(map[string]Job)}
}

// Put stores job and assigns the next id, "job-1", "job-2", and so on.
// A cancelled context returns a wrapped ctx.Err() and stores nothing.
func (s *MemoryStore) Put(ctx context.Context, job Job) (Job, error) {
	if ctx.Err() != nil {
		return Job{}, fmt.Errorf("context canceled: %w", ctx.Err())
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// check if the context is canceled
	select {
	case <-ctx.Done():
		return Job{}, fmt.Errorf("context canceled: %w", ctx.Err())
	default:
		// context is not canceled
		s.jobs[job.ID] = job
		s.next = len(s.jobs) + 1
		return job, nil
	}
}

// Get returns the job for id. A missing id wraps ErrNotFound.
// A cancelled context returns a wrapped ctx.Err().
func (s *MemoryStore) Get(ctx context.Context, id string) (Job, error) {
	if ctx.Err() != nil {
		return Job{}, fmt.Errorf("context canceled: %w", ctx.Err())
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return Job{}, fmt.Errorf("context canceled: %w", ctx.Err())
	default:
		job, ok := s.jobs[id]
		if !ok {
			return Job{}, ErrNotFound
		}
		return job, nil
	}
}

// NewMux serves POST /jobs and GET /jobs/{id}.
// An empty name or bad JSON is 400. A missing job is 404. Other store errors are 500.
func NewMux(store Store) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /jobs", func(w http.ResponseWriter, r *http.Request) {
		handlePost(w, r, store)
	})
	mux.HandleFunc("GET /jobs/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, store)
	})
	return mux
}

func handlePost(w http.ResponseWriter, r *http.Request, store Store) {
	job := Job{}
	err := json.NewDecoder(r.Body).Decode(&job)

	// fmt.Println("JobId", job.ID)
	// fmt.Println("JobName", job.Name)
	// fmt.Println("Job", job)

	if err != nil || job.ID == "" || job.Name == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	job, err = store.Put(r.Context(), job)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func handleGet(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.PathValue("id")
	job, err := store.Get(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(job)
}

type checks struct {
	failed bool
}

func (c *checks) report(name string, ok bool) {
	if ok {
		fmt.Printf("PASS %s\n", name)
		return
	}
	fmt.Printf("FAIL %s\n", name)
	c.failed = true
}

func main() {
	var c checks
	store := NewMemoryStore()
	mux := NewMux(store)

	req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{"id":"job-1", "name":"compile"}`))
	rec := httptest.NewRecorder()
	// fmt.Println("request", req.Method, req.URL.Path)
	mux.ServeHTTP(rec, req)
	// fmt.Println("response", rec.Code, rec.Body.String())
	var created Job
	decodeErr := json.NewDecoder(rec.Body).Decode(&created)
	c.report("post-code", rec.Code == http.StatusCreated)
	c.report("post-job", decodeErr == nil && created.Name == "compile" && strings.HasPrefix(created.ID, "job-"))
	c.report("post-id", created.ID == "job-1")

	req = httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{}`))
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	c.report("post-empty", rec.Code == http.StatusBadRequest && rec.Body.String() == "bad request\n")
	// c.report("post-empty", rec.Code == http.StatusBadRequest)

	req = httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`not json`))
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	c.report("post-bad-json", rec.Code == http.StatusBadRequest && rec.Body.String() == "bad request\n")

	req = httptest.NewRequest(http.MethodGet, "/jobs/"+created.ID, nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var fetched Job
	decodeErr = json.NewDecoder(rec.Body).Decode(&fetched)
	c.report("get-code", rec.Code == http.StatusOK)
	c.report("get-job", decodeErr == nil && fetched == created)

	req = httptest.NewRequest(http.MethodGet, "/jobs/missing", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	c.report("missing", rec.Code == http.StatusNotFound && rec.Body.String() == "not found\n")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := store.Get(ctx, created.ID)
	c.report("canceled", errors.Is(err, context.Canceled))
	_, err = store.Put(ctx, Job{Name: "later"})
	c.report("canceled-put", errors.Is(err, context.Canceled))

	if c.failed {
		os.Exit(1)
	}
}
