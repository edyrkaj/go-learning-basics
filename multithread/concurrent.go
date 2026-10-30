package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Signal to WaitGroup that this worker is finished when done
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second) // Simulates work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// Increment WaitGroup counter for each goroutine launched
		wg.Add(1)

		// Pass pointer to WaitGroup so worker modifies the original counter
		go worker(i, &wg)
	}

	// Block until WaitGroup counter reaches zero
	wg.Wait()

	// Total time will be ~1 second (all 3 workers ran concurrently)
	fmt.Printf("Concurrent execution took: %v\n", time.Since(start))
}
