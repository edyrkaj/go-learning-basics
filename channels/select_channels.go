package main

import (
	"fmt"
	"time"
)

func fetchData(ch chan string) {
	// Simulate a slow network request (3 seconds)
	time.Sleep(3 * time.Second)
	ch <- "Database record payload"
}

func main() {
	mode := 2 // 1 for select, 2 for non-blocking
	if mode == 1 {
		run_select()
	} else {
		run_non_blocking()
	}
}

func run_select() {
	ch := make(chan string)

	go fetchData(ch)

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // explanation: the timeout triggers after 1 second
		// Triggers after 1s because fetchData takes 3s
		fmt.Println("Error: Operation timed out!")
	}
}

// run_non_blocking demonstrates non-blocking receive/send and multiplexing across multiple channels
func run_non_blocking() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message available")
	}

	// Non-blocking send
	msg := "Hello"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No receiver ready, skipped sending")
	}

	// Multiplexing check across multiple channels
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity on any channel")
	}
}
