package calculator

import (
	"errors"
	"fmt"
)

// Add adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts b from a and returns the result
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers and returns the result
func Multiply(a, b int) int {
	return a * b
}

// Divide divides a by b and returns the result or an error if b is zero
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return float64(a) / float64(b), nil
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd checks if a number is odd
func IsOdd(n int) bool {
	return n%2 != 0
}

// CalculateFactorial calculates the factorial of a non-negative integer
func CalculateFactorial(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial not defined for negative numbers")
	}
	if n > 20 {
		return 0, fmt.Errorf("number too large, would cause overflow")
	}

	var result uint64 = 1
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result, nil
}

// GCD calculates the Greatest Common Divisor of two integers
func GCD(a, b int) int {
	// Convert negative numbers to positive
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
