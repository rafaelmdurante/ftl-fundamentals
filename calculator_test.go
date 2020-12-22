package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b, want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{a: -5, b: -5, want: -10},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 5, b: 3, want: 2},
		{a: 10, b: 5, want: 5},
		{a: 10, b: -5, want: 15},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 5, b: 3, want: 15},
		{a: 10, b: 5, want: 50},
		{a: 2, b: -2, want: -4},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b, want  float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 6, b: 0, want: 0, errExpected: true},
		{a: 6, b: 3, want: 2, errExpected: false},
		{a: 3, b: 2, want: 1.5, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%f, %f): unexpected error: %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
