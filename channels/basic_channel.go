package main

import (
	"fmt"
	"time"
)

func fetchResult(ch chan string) {
	time.Sleep(1 * time.Second) // Simulate async work

	// Send data into channel
	ch <- "Data successfully loaded from API"
}

func main() {
	// 1. Create an unbuffered channel of type string
	resultChan := make(chan string)

	// 2. Launch worker goroutine
	go fetchResult(resultChan)

	fmt.Println("Waiting for result...")

	// 3. Receive data from channel (blocks until data arrives)
	msg := <-resultChan

	fmt.Println("Received:", msg)
}
