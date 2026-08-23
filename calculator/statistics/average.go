package statistics

import (
	"errors"
	"math"
	"sort"
)

// CalculateAverage calculates the average (mean) of a slice of numbers
func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

// CalculateMedian calculates the median of a slice of numbers
func CalculateMedian(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate median of empty slice")
	}

	// Create a copy to avoid modifying the original slice
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	// If the length is odd, return the middle element
	if len(sorted)%2 == 1 {
		return sorted[len(sorted)/2], nil
	}

	// If the length is even, return the average of the two middle elements
	middle1 := sorted[len(sorted)/2-1]
	middle2 := sorted[len(sorted)/2]
	return (middle1 + middle2) / 2, nil
}

// CalculateVariance calculates the variance of a slice of numbers
func CalculateVariance(numbers []float64) (float64, error) {
	if len(numbers) <= 1 {
		return 0, errors.New("variance requires at least two data points")
	}

	mean := CalculateAverage(numbers)
	sum := 0.0

	for _, num := range numbers {
		difference := num - mean
		sum += difference * difference
	}

	// Using sample variance (n-1 denominator)
	return sum / float64(len(numbers)-1), nil
}

// CalculateStandardDeviation calculates the standard deviation of a slice of numbers
func CalculateStandardDeviation(numbers []float64) (float64, error) {
	variance, err := CalculateVariance(numbers)
	if err != nil {
		return 0, err
	}

	return math.Sqrt(variance), nil
}

// CalculateMode calculates the mode (most frequent value) of a slice of numbers
func CalculateMode(numbers []float64) ([]float64, error) {
	if len(numbers) == 0 {
		return nil, errors.New("cannot calculate mode of empty slice")
	}

	// Count occurrences of each number
	counts := make(map[float64]int)
	for _, num := range numbers {
		counts[num]++
	}

	// Find the highest frequency
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Find all values with the highest frequency
	modes := []float64{}
	for num, count := range counts {
		if count == maxCount {
			modes = append(modes, num)
		}
	}

	return modes, nil
}
