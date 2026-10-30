package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementCounter(counter *atomic.Int64) {
	fmt.Println("incrementing counter ", counter.Load())
	counter.Add(1)
}

func main() {
	var wg sync.WaitGroup

	// Declare an atomic int64 counter (initial value 0)
	var counter atomic.Int64

	// Launch 1,000 goroutines concurrently incrementing the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			incrementCounter(&counter)
		}()
	}

	wg.Wait()

	// Read counter value safely
	fmt.Printf("Final Counter Value: %d\n", counter.Load())
}

/**
* Feature	sync.Mutex	sync/atomic
Scope	Complex operations, multi-line critical sections, maps, slices, structs.	Single primitive variables (counters, booleans, pointers).
Mechanics	Blocks execution; context switches threads if contention is high.	Non-blocking CPU hardware instruction.
Performance	High throughput, slight locking overhead under contention.	Ultra-high performance for simple state checks and increments.
Safety Risk	Deadlocks, forgetting to unlock.	Subtle race conditions if combining multiple atomic variables together.
*/
