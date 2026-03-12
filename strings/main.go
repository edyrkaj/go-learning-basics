package main

import (
	"fmt"
	"slices"
)

func main() {
	given_word := "Hello, World!"
	reversed_word := reverse(given_word)
	fmt.Println(reversed_word)
}

func reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
