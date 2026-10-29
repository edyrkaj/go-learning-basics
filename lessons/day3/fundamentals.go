package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

// DedupInPlace keeps the first occurrence of each value, in that order.
// The result shares the caller's backing array. Empty and nil inputs return
// an empty non-nil slice.
func DedupInPlace(nums []int) []int {
	return nil
}

// WordCount counts words split on Unicode whitespace. Words are lowercased.
// Characters that are not letters or digits are removed from both ends.
// Interior punctuation stays. The result is never nil.
func WordCount(text string) map[string]int {
	return nil
}

// Greeter is satisfied by any value with a Greet method.
type Greeter interface {
	Greet() string
}

// Person uses a value receiver, so both Person and *Person satisfy Greeter.
type Person struct {
	Name string
}

func (p Person) Greet() string {
	return ""
}

// Adder is satisfied only by *Counter, because the methods use pointer receivers.
// A Counter value cannot be assigned to an Adder.
type Adder interface {
	Add(n int)
	Value() int
}

// Counter is a running total. Methods use pointer receivers so Add can mutate it.
type Counter struct {
	n int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(n int) {}

func (c *Counter) Value() int {
	return 0
}

// ErrNotFound is the sentinel returned when FindUser is given a non-positive id.
var ErrNotFound = errors.New("not found")

// FindUser returns "user-<id>" for a positive id.
// A non-positive id returns a wrapped ErrNotFound.
func FindUser(id int) (string, error) {
	return "", nil
}

// IsNotFound reports whether err wraps ErrNotFound.
func IsNotFound(err error) bool {
	return false
}

// DeferOrder records "a", then "b", then "c" with defer.
// Deferred calls run last-in first-out. The named result lets those calls
// append to the slice that the caller receives: "c", "b", "a".
func DeferOrder() (order []string) {
	return nil
}

// SafeDiv returns a/b. Division by zero panics and is recovered here.
// Named results let the deferred recover set the returned values.
func SafeDiv(a, b int) (result int, recovered bool) {
	return 0, false
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

	c.report("dedup", slices.Equal(DedupInPlace([]int{1, 2, 1, 3, 2}), []int{1, 2, 3}))
	empty := DedupInPlace(nil)
	c.report("dedup-empty", empty != nil && len(empty) == 0)

	got := WordCount("Go, go!")
	c.report("wordcount", got != nil && len(got) == 1 && got["go"] == 2)

	var g Greeter = Person{Name: "Ada"}
	c.report("greet-value", g.Greet() == "hi Ada")
	g = &Person{Name: "Ada"}
	c.report("greet-pointer", g.Greet() == "hi Ada")

	var adder Adder = NewCounter()
	adder.Add(2)
	adder.Add(3)
	c.report("counter", adder.Value() == 5)

	_, err := FindUser(0)
	c.report("find-missing", IsNotFound(err))
	name, err := FindUser(4)
	c.report("find-user", err == nil && name == "user-4")

	c.report("defer-order", slices.Equal(DeferOrder(), []string{"c", "b", "a"}))

	q, recovered := SafeDiv(8, 2)
	c.report("safe-div", q == 4 && !recovered)
	q, recovered = SafeDiv(1, 0)
	c.report("safe-div-zero", q == 0 && recovered)

	if c.failed {
		os.Exit(1)
	}
}
