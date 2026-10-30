// string_test.go
package main

import (
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	// b.N is dynamically adjusted by the Go runner until it gets a stable timing sample
	for i := 0; i < b.N; i++ {
		_ = "hello" + "world"
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ReportAllocs() // Tracks memory allocation counts and byte size

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString("hello")
		builder.WriteString("world")
		_ = builder.String()
	}
}

// go test -bench=. string_test.go
// go test -bench=, -memprofile=mem.out
// go test -bench=, -cpuprofile=cpu.out
// go tool pprof -http=:8080 mem.out
