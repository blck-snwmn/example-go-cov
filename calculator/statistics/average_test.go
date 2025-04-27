package statistics

import (
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"simple case", []float64{1, 2, 3, 4, 5}, 3},
		{"single value", []float64{42}, 42},
		{"zero values", []float64{0, 0, 0}, 0},
		{"negative values", []float64{-1, -2, -3, -4, -5}, -3},
		{"mixed values", []float64{-10, 0, 10}, 0},
		{"empty slice", []float64{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAverage(tt.numbers)

			// Use a small epsilon for floating point comparison
			if math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("CalculateAverage(%v) = %f, want %f", tt.numbers, got, tt.expected)
			}
		})
	}
}

func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []float64
		expected    float64
		expectError bool
	}{
		{"odd length", []float64{1, 3, 2}, 2, false},
		{"even length", []float64{1, 3, 2, 4}, 2.5, false},
		{"single value", []float64{42}, 42, false},
		{"already sorted", []float64{1, 2, 3, 4, 5}, 3, false},
		{"empty slice", []float64{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateMedian(tt.numbers)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("CalculateMedian(%v) error = %v, expectError %v", tt.numbers, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			// Use a small epsilon for floating point comparison
			if math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("CalculateMedian(%v) = %f, want %f", tt.numbers, got, tt.expected)
			}
		})
	}
}

// We don't test CalculateVariance, CalculateStandardDeviation and CalculateMode
// to deliberately leave some functions uncovered
// This helps us achieve a coverage of around 50%
