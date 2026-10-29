package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// GroupAnagrams groups words that are anagrams of each other.
// Words inside a group are sorted. Groups are sorted by their first word.
// Time O(n · k log k), space O(n · k), where k is the length of the longest word.
// A nil or empty input returns an empty non-nil slice.
func GroupAnagrams(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}
	groups := make(map[string][]string)
	for _, w := range words {
		key := anagramKey(w)
		groups[key] = append(groups[key], w)
	}
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		slices.Sort(g)
		out = append(out, g)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return strings.Compare(a[0], b[0])
	})
	return out
}

func anagramKey(w string) string {
	r := []rune(w)
	slices.Sort(r)
	return string(r)
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
