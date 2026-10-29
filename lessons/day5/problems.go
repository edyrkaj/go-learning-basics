package main

import (
	"fmt"
	"os"
	"sync"
)

// TwoSum returns the indexes of the pair that sums to target.
// The pair is the one whose second index is smallest.
// One map pass: time O(n), space O(n).
func TwoSum(nums []int, target int) (i, j int, ok bool) {
	return 0, 0, false
}

// ValidParentheses reports whether brackets (), [], and {} are balanced.
// Any other rune is invalid. Time O(n), space O(n).
func ValidParentheses(s string) bool {
	return false
}

// SafeCounter is safe for concurrent Add and Value calls.
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

// Add changes the counter by delta.
func (c *SafeCounter) Add(delta int) {}

// Value returns the current total.
func (c *SafeCounter) Value() int {
	return 0
}

type checks struct {
	failed bool
}

func (c *checks) report(name string, ok bool) {
	if ok {
		fmt.Printf("PASS %s\n", name)
		return
	}
	fmt.Printf("FAIL %s\n", name)
	c.failed = true
}

func main() {
	var c checks

	i, j, ok := TwoSum([]int{2, 7, 11, 15}, 9)
	c.report("two-sum", ok && i == 0 && j == 1)
	_, _, ok = TwoSum([]int{1, 2, 3}, 100)
	c.report("two-sum-miss", !ok)

	c.report("parens-ok", ValidParentheses("()[]{}"))
	c.report("parens-empty", ValidParentheses(""))
	c.report("parens-bad", !ValidParentheses("(]"))
	c.report("parens-open", !ValidParentheses("("))

	var counter SafeCounter
	var wg sync.WaitGroup
	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	c.report("counter", counter.Value() == 100)

	if c.failed {
		os.Exit(1)
	}
}
