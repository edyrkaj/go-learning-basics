package main

import (
	"fmt"
	"slices"
)

func main() {
	given_word := "Hello, World!"
	reversed_word := Reverse(given_word)
	fmt.Println(reversed_word)

	old_reversed_word := Old_Reverse(given_word)
	fmt.Println(old_reversed_word)
	// compare the two functions

	if slices.Equal([]rune(reversed_word), []rune(old_reversed_word)) {
		fmt.Println("The two functions are equal")
	} else {
		fmt.Println("The two functions are not equal")
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func Old_Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
