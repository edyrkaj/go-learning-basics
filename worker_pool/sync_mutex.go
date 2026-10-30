package main

import (
	"fmt"
	"sync"
)

// SafeCounter wraps a shared variable alongside its protective mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc safely increments the counter
func (c *SafeCounter) Inc() {
	c.mu.Lock()         // Acquire lock: no other goroutine can enter this block
	defer c.mu.Unlock() // Release lock when function exits

	fmt.Printf("Incrementing counter... [%d] -> ", c.value)
	c.value++
	fmt.Printf("%d\n", c.value)
}

// Value safely retrieves the counter's current value
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := SafeCounter{}

	// Launch 1,000 goroutines concurrently incrementing the same counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	// Guaranteed to output 1000 due to mutex protection
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
