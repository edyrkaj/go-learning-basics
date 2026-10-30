package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, limiter <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Wait for the rate-limiter pulse before processing
		<-limiter

		fmt.Printf("[%s] Worker %d started processing Job %d\n",
			time.Now().Format("15:04:05.000"), id, job.ID)

		// Simulate job task
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// Set rate limit: 1 job every 200ms (5 jobs/sec)
	rate := 200 * time.Millisecond
	ticker := time.NewTicker(rate)
	defer ticker.Stop() // Always stop ticker to prevent leaks

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// 1. Launch worker pool passing ticker.C as the rate limiter
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, ticker.C, &wg)
	}

	// 2. Queue jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	// 3. Wait for workers to finish
	wg.Wait()
	fmt.Println("\nAll jobs processed within rate limit constraints.")
}
