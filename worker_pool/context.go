package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Duration time.Duration
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		// 1. Check if context has been canceled or timed out
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Stopping due to context signal: %v\n", id, ctx.Err())
			return

		// 2. Fetch job from queue if available
		case job, ok := <-jobs:
			if !ok {
				// Jobs channel closed
				return
			}

			fmt.Printf("[Worker %d] Starting job %d (requires %v)\n", id, job.ID, job.Duration)

			// Simulate work that respects context cancellation during execution
			select {
			case <-time.After(job.Duration):
				fmt.Printf("[Worker %d] Completed job %d\n", id, job.ID)
			case <-ctx.Done():
				fmt.Printf("[Worker %d] Aborted job %d mid-task: %v\n", id, job.ID, ctx.Err())
				return
			}
		}
	}
}

func main() {
	// Set a deadline of 2 seconds for all work to complete
	timeoutSec := 2
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel() // Always call cancel to release resources if work finishes early

	const numJobs = 5
	const numWorkers = 2

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Launch workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	// Submit jobs (some short, some long)
	jobsList := []Job{
		{ID: 1, Duration: 500 * time.Millisecond},
		{ID: 2, Duration: 800 * time.Millisecond},
		{ID: 3, Duration: 1500 * time.Millisecond}, // Will trigger timeout before completing // explanation:
		{ID: 4, Duration: 500 * time.Millisecond},
		{ID: 5, Duration: 500 * time.Millisecond},
	}

	for _, j := range jobsList {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish exiting
	wg.Wait()
	fmt.Println("Main: All workers shut down cleanly.")
}
