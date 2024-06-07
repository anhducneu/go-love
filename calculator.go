package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a float64, b int64) float64 {
	return a - float64(b)
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Devine(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot devine by zero")
	}

	return a / b, nil
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}

	return math.Sqrt(x), nil
}
