package testing

import (
	"testing"
)

// ----------------------------------------------------------------------------
// 1. UNIT TEST EXAMPLE
// Run with: go test -v -run TestAdd
// ----------------------------------------------------------------------------
func TestAdd(t *testing.T) {
	// Simple unit test verifying Add function behaves correctly
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}
}

// Table-driven unit test example for more thorough testing
// Run with: go test -v -run TestDivide
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"valid division", 10, 2, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative result", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.expectError {
				t.Errorf("Divide() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if result != tt.expected {
				t.Errorf("Divide() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// ----------------------------------------------------------------------------
// 2. INTEGRATION TEST EXAMPLE
// Run with: go test -v -run TestProcessDataIntegration
// ----------------------------------------------------------------------------
// Integration tests often check interactions between multiple components
// or external dependencies. We simulate this using `ProcessData`.
func TestProcessDataIntegration(t *testing.T) {
	// Optional: Skip integration tests using the -short flag (go test -short)
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This tests the integration between both Add and Divide internally
	// in the ProcessData func
	result, err := ProcessData(10) // (10 + 10) / 2 = 10
	if err != nil {
		t.Fatalf("ProcessData returned unexpected error: %v", err)
	}

	if result != 10 {
		t.Errorf("ProcessData(10) = %v, expected 10", result)
	}
}

// ----------------------------------------------------------------------------
// 3. BENCHMARK TEST EXAMPLE
// Run with: go test -bench=. -benchmem
// ----------------------------------------------------------------------------
func BenchmarkAdd(b *testing.B) {
	// The loop will run b.N times, where b.N is dynamically adjusted by the
	// benchmark runner until the timing results are stable.
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// ----------------------------------------------------------------------------
// 4. FUZZ TEST EXAMPLE (Go 1.18+)
// Run with: go test -fuzz=FuzzAdd -fuzztime=10s
// ----------------------------------------------------------------------------
func FuzzAdd(f *testing.F) {
	// Add seed corpus - baseline examples
	f.Add(1, 2)
	f.Add(5, -5)
	f.Add(100, 200)

	// Fuzz loop provides random values into the test case
	f.Fuzz(func(t *testing.T, a int, b int) {
		// Example check: Addition should be symmetric (a+b == b+a)
		result1 := Add(a, b)
		result2 := Add(b, a)

		if result1 != result2 {
			t.Errorf("Addition is not symmetric for %d and %d: %d != %d", a, b, result1, result2)
		}
	})
}
