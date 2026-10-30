package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the task data
type Job struct {
	ID int
}

// Result represents the output of processing a Job
type Result struct {
	Job    Job
	Output string
}

// worker function runs as a fixed goroutine in the pool
// jobs <-chan Job: receives jobs from the jobs channel
// results chan<- Result: sends results to the results channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Worker processes jobs until the 'jobs' channel is closed
	for job := range jobs {
		fmt.Printf("[Worker %d] Started job %d\n", id, job.ID)

		// Simulate dynamic work duration
		time.Sleep(500 * time.Millisecond)

		output := fmt.Sprintf("Result of job %d processed by worker %d", job.ID, id)
		fmt.Printf("[Worker %d] Finished job %d\n", id, job.ID)

		// Send completed result to results channel
		results <- Result{Job: job, Output: output}
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// Create buffered channels for jobs and results
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// 1. Start fixed pool of workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	// Close jobs channel to signal workers no more incoming work
	close(jobs)

	// 3. Wait for all workers to complete in a separate goroutine, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Collect results as they arrive
	for result := range results {
		fmt.Printf("Main: Received -> %s\n", result.Output)
	}

	fmt.Println("All work completed successfully.")
}

// 1. issues to be considered when working on worker pools
/**
 * Common Leak Causes & Fixes:
 *
 * Unbuffered/Full Channel Blocks: If a worker sends a result to results <- res but the main thread stops reading (or crashes), the worker blocks forever on that line.
 *
 * Fix: Ensure channel buffers match job counts (make(chan Result, numJobs)), or use context.Context to handle early cancellation.
 *
 * Leaking Worker Goroutines: Forgetting to close the jobs channel keeps workers waiting on range jobs indefinitely.
 *
 * Fix: Always close(jobs) when done dispatching work, and use defer wg.Done() to guarantee workers deregister even during a crash.
 *
 * Context Cancellation: Use Go's context.Context so workers stop processing immediately when the overall application receives a shutdown signal or a global deadline expires.
 */

// 2. Robust error handling
/**
* In Go, errors are values. To handle errors gracefully in a worker pool without crashing the pool, pass errors out through your Result struct or use specialized sync primitives.

Pattern A: Returning Errors via Result Struct
Embed an error field inside your result object so the caller decides how to handle failures without killing the worker.
*/

// 3. Use Recover
// Use recover() in worker goroutines to handle panics gracefully and avoid crashing the worker pool.
//
// User accepted prediction:
//

// 4. Graceful Shutdown
// Use a context.Context to gracefully shut down the worker pool, ensuring all workers exit cleanly.
//
// User accepted prediction:
