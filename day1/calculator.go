package calculator

import (
	"errors"
	"math"
)

var WrongInputError = "Wrong Input Error"

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Minus(num1, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) float64 {
	return num1 / num2
}

func Sin(angle float64) float64 {
	return math.Sin(angle)
}

func Cos(angle float64) float64 {
	return math.Cos(angle)
}

func Tan(angle float64) float64 {
	return math.Tan(angle)
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New(WrongInputError)
	}
	return math.Sqrt(num), nil
}
