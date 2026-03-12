package testing

import (
	"errors"
)

// Add sums two integers
func Add(a, b int) int {
	return a + b
}

// Divide divides a by b, returning an error if b is 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// ProcessData simulates a complex operation integrating multiple steps
func ProcessData(input int) (int, error) {
	res1 := Add(input, 10)
	res2, err := Divide(res1, 2)
	if err != nil {
		return 0, err
	}
	return res2, nil
}
