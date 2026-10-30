package main

import "fmt"

func sumEvenNumbers(limit int) (sum int) {
	sum = 0

	// start loop
	for i := 1; i <= limit; i++ {
		// verify the even number
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	limit := 10
	result := sumEvenNumbers(limit)
	// Expected output for 10 should be 30 (2 + 4 + 6 + 8 + 10)
	fmt.Printf("Sum of even numbers up to %d is: %d\n", limit, result)
}
