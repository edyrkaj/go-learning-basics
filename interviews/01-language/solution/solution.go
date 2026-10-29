package main

import (
	"fmt"
	"os"
	"slices"
)

// Store is a named value. A nil Store means there is no value.
type Store interface {
	Get() string
}

// Mem is an in-memory Store. Get uses a pointer receiver.
type Mem struct {
	Name string
}

func (m *Mem) Get() string {
	return m.Name
}

// Prefix returns the first n elements in a new slice.
// n is in range. Appending to the result must not change nums.
func Prefix(nums []int, n int) []int {
	out := make([]int, n)
	copy(out, nums[:n])
	return out
}

// Label returns the value for key and whether the key was present.
func Label(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}

// NewStore returns a Store for name. An empty name returns a nil interface,
// not a nil *Mem stored inside a non-nil interface.
func NewStore(name string) Store {
	if name == "" {
		return nil
	}
	return &Mem{Name: name}
}

// IsEmpty reports whether store is a nil interface.
func IsEmpty(store Store) bool {
	return store == nil
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

	nums := []int{1, 2, 3, 4}
	got := Prefix(nums, 3)
	c.report("prefix-values", slices.Equal(got, []int{1, 2, 3}))
	got = append(got, 99)
	c.report("prefix-isolated", slices.Equal(nums, []int{1, 2, 3, 4}))

	v, ok := Label(map[string]int{"zero": 0}, "zero")
	c.report("label-zero", v == 0 && ok)
	_, ok = Label(map[string]int{}, "missing")
	c.report("label-missing", !ok)

	c.report("store-empty", IsEmpty(NewStore("")))
	c.report("store-value", NewStore("mem").Get() == "mem")

	if c.failed {
		os.Exit(1)
	}
}
