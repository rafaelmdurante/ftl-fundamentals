package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
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

func TestAddRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if got != want {
			t.Fatalf("AddRandom(%f, %f): want %f, got %f", a, b, want, got)
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
		{a: 6, b: 0, want: 999, errExpected: true},
		{a: 6, b: 3, want: 2, errExpected: false},
		{a: 3, b: 2, want: 1.5, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			// another syntax: if tc.errExpected != (err != nil) { t.Fatalf(...) }
			// t.Fatal is used instead of t.Errorf because we want to exit after printing msg
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		num, want   float64
		errExpected bool
	}
	testCases := []testCase{
		{num: 9, want: 3, errExpected: false},
		{num: 100, want: 10, errExpected: false},
		{num: -100, want: 666, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.num)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Sqrt(%f): unexpected error status: %v", tc.num, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.num, tc.want, got)
		}
	}
}

type manyArgsTestCase struct {
	inputs []float64
	want   float64
}

func TestAddMany(t *testing.T) {
	t.Parallel()
	testCases := []manyArgsTestCase{
		{inputs: []float64{1, 2}, want: 3},
		{inputs: []float64{1, 2, 3}, want: 6},
		{inputs: []float64{1, 2, 3, 4, 5}, want: 15},
	}
	for _, tc := range testCases {
		got := calculator.AddMany(tc.inputs...)
		if got != tc.want {
			t.Errorf("AddMany(%v): want: %f, got: %f", tc.inputs, tc.want, got)
		}
	}
}

func TestSubtractMany(t *testing.T) {
	t.Parallel()
	testCases := []manyArgsTestCase{
		{inputs: []float64{8, 2}, want: 6},
		{inputs: []float64{7, 2, 3}, want: 2},
		{inputs: []float64{10, 2, 3, 4, 5}, want: -4},
	}
	for _, tc := range testCases {
		got := calculator.SubtractMany(tc.inputs...)
		if got != tc.want {
			t.Errorf("SubtractMany(%v): want: %f, got: %f", tc.inputs, tc.want, got)
		}
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input string
		want  float64
	}
	testCases := []testCase{
		{input: "2 * 2", want: 4},
		{input: "1 + 1.5", want: 2.5},
		{input: "18  /   6", want: 3},
		{input: "100 - 0.1", want: 99.9},
	}
	for _, tc := range testCases {
		got := calculator.Evaluate(tc.input)
		if got != tc.want {
			t.Errorf("Evaluate(%s): want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
