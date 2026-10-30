package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i * 10
		time.Sleep(500 * time.Millisecond)
	}

	close(ch) // explanation: the channel is closed to signal the end of production
}

func main() {
	// Buffered channel with capacity 3
	ch := make(chan int, 3) // explanation: a buffered channel with capacity 3

	go producer(ch) // explanation: the producer goroutine sends values to the channel

	// Loop receives values until 'ch' is closed and drained
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}

	fmt.Println("Channel closed, iteration finished.") // explanation: the loop is exited when the channel is closed
}
