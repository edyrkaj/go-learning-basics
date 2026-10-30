package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for range 10 {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for n := range 10 {
		fmt.Printf("Step %d \n", n)
		<-done
	}
}
