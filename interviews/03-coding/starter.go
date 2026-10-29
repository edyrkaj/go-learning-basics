package main

import (
	"fmt"
	"os"
	"slices"
)

// GroupAnagrams groups words that are anagrams of each other.
// Words inside a group are sorted. Groups are sorted by their first word.
// A nil or empty input returns an empty non-nil slice.
func GroupAnagrams(words []string) [][]string {
	return nil
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

func sameGroups(got, want [][]string) bool {
	if got == nil || len(got) != len(want) {
		return false
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			return false
		}
	}
	return true
}

func main() {
	var c checks
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}}
	c.report("groups", sameGroups(got, want))
	empty := GroupAnagrams(nil)
	c.report("groups-empty", empty != nil && len(empty) == 0)
	if c.failed {
		os.Exit(1)
	}
}
