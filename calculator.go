// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing them together
// If b is zero, it also returns an error message
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

// Sqrt takes a number and returns its square root
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("bad input: %f (no real number squared equals a negative)", num)

	}
	return math.Sqrt(num), nil
}
