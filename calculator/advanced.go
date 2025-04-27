package calculator

import (
	"errors"
	"math"
)

// Power calculates a raised to the power of b
func Power(a, b float64) (float64, error) {
	if a == 0 && b == 0 {
		return 0, errors.New("0^0 is undefined")
	}

	if a < 0 && math.Floor(b) != b {
		return 0, errors.New("negative base with non-integer exponent is not a real number")
	}

	return math.Pow(a, b), nil
}

// Sqrt calculates the square root of a number
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number is not a real number")
	}
	return math.Sqrt(a), nil
}

// Log calculates the natural logarithm of a number
func Log(a float64) (float64, error) {
	if a <= 0 {
		return 0, errors.New("logarithm of non-positive number is undefined")
	}
	return math.Log(a), nil
}

// Log10 calculates the base-10 logarithm of a number
func Log10(a float64) (float64, error) {
	if a <= 0 {
		return 0, errors.New("logarithm of non-positive number is undefined")
	}
	return math.Log10(a), nil
}

// Round rounds a number to the nearest integer
func Round(a float64) float64 {
	return math.Round(a)
}

// Floor returns the greatest integer less than or equal to a
func Floor(a float64) float64 {
	return math.Floor(a)
}

// Ceiling returns the least integer greater than or equal to a
func Ceiling(a float64) float64 {
	return math.Ceil(a)
}
