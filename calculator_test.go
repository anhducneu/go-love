package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	var testCases = []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 3},
		{a: 1, b: 0, want: 1},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Want %f, got %f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()

	var want float64 = 1.5
	got := calculator.Subtract(2.5, 1.0)
	if want != got {
		t.Errorf("Want %f, got %f", want, got)
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()

	var want float64 = 6
	got := calculator.Multiply(2, 3)
	if want != got {
		t.Errorf("Want %f, got %f", want, got)
	}
}

func TestDevine(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	var testCases = []testCase{
		{2, 2, 1},
		{4, 2, 2},
		{15, 3, 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Devine(tc.a, tc.b)

		if err != nil {
			t.Fatalf("Devine(%f, %f) returned an error: %v", tc.a, tc.b, err)
		}

		if !CloseEnough(tc.want, got, 0.0001) {
			t.Errorf("Want %f, got %f", tc.want, got)
		}
	}
}

func TestDevineInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Devine(1, 0)

	if err == nil {
		t.Error("Devine(1, 0) should return an error but it didn't")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		x    float64
		want float64
	}
	var testCases = []testCase{
		{x: 4, want: 2},
		{x: 9, want: 3},
		{x: 16, want: 4},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.x)

		if err != nil {
			t.Fatalf("Sqrt(%f) returned an error: %v", tc.x, err)
		}

		if !CloseEnough(tc.want, got, 0.1) {
			t.Errorf("Want %f, got %f", tc.want, got)
		}

	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)

	if err == nil {
		t.Error("Sqrt(-1) should return an error but it didn't")
	}

}

func CloseEnough(a, b, tolerance float64) bool {
	return a == b || math.Abs(a-b) < tolerance
}

func TestDataType(t *testing.T) {

	var x int64 = 100
	y := x
	y = 200

	fmt.Println(x)
	fmt.Println(y)

}
