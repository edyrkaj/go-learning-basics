package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

// User embeds Logger (Composition)
type User struct {
	Logger // Embedded field: User inherits Logger's methods directly
	Name   string
}

func main() {
	u := User{Name: "Alice"}

	// Call Log directly on User thanks to method promotion
	u.Log("User account accessed") 
}