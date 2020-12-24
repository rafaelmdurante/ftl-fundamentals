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

// AddMany takes all numbers and returns the result of adding them together.
func AddMany(inputs ...float64) float64 {
	sum := 0.
	for _, v := range inputs {
		sum += v
	}
	return sum
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// SubtractMany takes all numbers and returns the result of subtracting them.
func SubtractMany(inputs ...float64) float64 {
	// store the first input
	result := inputs[0]
	// iterate from input 1 to the end
	for _, v := range inputs[1:] {
		result -= v
	}
	return result
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

// Evaluate takes a string of arithmetic operations and returns its result
// It accepts only two arguments in the arithmetic operation.
// It handles addition, subtraction, multiplication and division.
func Evaluate(input string) float64 {
	var a, b float64
	var operator string

	_, err := fmt.Sscanf(input, "%f %s %f", &a, &operator, &b)
	if err != nil {
		panic(err)
	}

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	return result
}
