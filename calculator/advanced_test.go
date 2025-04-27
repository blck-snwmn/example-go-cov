package calculator

import (
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{"simple case", 2, 3, 8, false},
		{"negative base, integer exponent", -2, 3, -8, false},
		{"zero base, non-zero exponent", 0, 5, 0, false},
		{"non-zero base, zero exponent", 5, 0, 1, false},
		{"negative exponent", 2, -2, 0.25, false},
		{"zero base, zero exponent", 0, 0, 0, true},
		{"negative base, non-integer exponent", -2, 0.5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.a, tt.b)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Power(%f, %f) error = %v, expectError %v", tt.a, tt.b, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			// Use a small epsilon for floating point comparison
			if math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Power(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		expected    float64
		expectError bool
	}{
		{"perfect square", 16, 4, false},
		{"non-perfect square", 10, math.Sqrt(10), false},
		{"zero", 0, 0, false},
		{"negative number", -4, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sqrt(tt.a)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Sqrt(%f) error = %v, expectError %v", tt.a, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			// Use a small epsilon for floating point comparison
			if math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Sqrt(%f) = %f, want %f", tt.a, got, tt.expected)
			}
		})
	}
}

// We don't test Log, Log10, Round, Floor, and Ceiling to deliberately leave some functions uncovered
// This helps us achieve a coverage of around 50%
