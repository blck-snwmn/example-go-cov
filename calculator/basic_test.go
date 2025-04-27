package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", -5, 8, 3},
		{"zero values", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 8, 3, 5},
		{"negative numbers", -8, -3, -5},
		{"mixed signs", -5, 3, -8},
		{"zero values", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); got != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 4, 3, 12},
		{"negative numbers", -4, -3, 12},
		{"mixed signs", -4, 3, -12},
		{"zero values", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); got != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    float64
		expectError bool
	}{
		{"positive numbers", 6, 3, 2.0, false},
		{"negative numbers", -6, -3, 2.0, false},
		{"mixed signs", -6, 3, -2.0, false},
		{"fractional result", 5, 2, 2.5, false},
		{"divide by zero", 5, 0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Divide(%d, %d) error = %v, expectError %v", tt.a, tt.b, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if got != tt.expected {
				t.Errorf("Divide(%d, %d) = %f, want %f", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// We don't test IsEven and IsOdd to deliberately leave some functions uncovered
// This helps us achieve a coverage of around 50%

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"simple case", 12, 8, 4},
		{"coprime numbers", 7, 5, 1},
		{"one number is multiple of other", 15, 5, 5},
		{"negative numbers", -12, 8, 4}, // GCD should handle negative numbers
		{"zero and non-zero", 0, 5, 5},  // GCD(0, n) = n
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.a, tt.b); got != tt.expected {
				t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
