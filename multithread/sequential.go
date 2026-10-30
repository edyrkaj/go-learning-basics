package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second) // Simulates work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	start := time.Now()

	// Runs workers one by one
	for i := 1; i <= 3; i++ {
		worker(i)
	}

	// Total time will be ~3 seconds
	fmt.Printf("Sequential execution took: %v\n", time.Since(start))
}
