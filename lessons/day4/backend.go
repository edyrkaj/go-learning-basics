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
	return Job{}, nil
}

// Get returns the job for id. A missing id wraps ErrNotFound.
// A cancelled context returns a wrapped ctx.Err().
func (s *MemoryStore) Get(ctx context.Context, id string) (Job, error) {
	return Job{}, nil
}

// NewMux serves POST /jobs and GET /jobs/{id}.
// An empty name or bad JSON is 400. A missing job is 404. Other store errors are 500.
func NewMux(store Store) http.Handler {
	return http.NewServeMux()
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

	req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{"name":"compile"}`))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var created Job
	decodeErr := json.NewDecoder(rec.Body).Decode(&created)
	c.report("post-code", rec.Code == http.StatusCreated)
	c.report("post-job", decodeErr == nil && created.Name == "compile" && strings.HasPrefix(created.ID, "job-"))

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

	if c.failed {
		os.Exit(1)
	}
}
