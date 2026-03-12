package main

import (
	"fmt"
	"sync"
	"time"
)

// add example of channels and go routines running in parallel
// add example of channels and go routines running in parallel with buffer
// make a bit complex example of channels and go routines running in parallel
// add example of channels and go routines running in parallel with buffer and select
func main() {
	// runWithBuffer()
	// runComplexExample()
	// runWithBufferAndSelect()
	runWithSync()
}

func runWithBuffer() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

func runComplexExample() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for j := range jobs {
				fmt.Printf("worker %d processing job %d\n", id, j)
				time.Sleep(10 * time.Millisecond) // Simulate some work
				results <- j * 2
			}
		}(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect all results
	for a := 1; a <= 5; a++ {
		<-results
	}
}

func runWithBufferAndSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		ch1 <- "message from ch1"
	}()

	go func() {
		ch2 <- "message from ch2"
	}()

	// time.Sleep(1 * time.Second) // without this sleep, the program will exit before the select statement can read the messages

	// Read both messages
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received via select:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received via select:", msg2)
		}
	}
}

func runWithSync() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// wg.Add(len(array)) // without this line we will have a deadlock

	go func() {
		for _, value := range array {
			ch <- &value
		}

		close(ch) // without this line we will have a deadlock when we have wg.Add(1)
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // what will be printed here?
			// wg.Done() // when we have wg.Add(len(array)) we need to call wg.Done() for each iteration
		}

		wg.Done()
	}()

	wg.Wait()
}
