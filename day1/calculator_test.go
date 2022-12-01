package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Error("wrong answer")
	}
}

func TestMinus(t *testing.T) {
	result := Minus(1, 2)
	if result != -1 {
		t.Error("wrong answer")
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(12, 2)
	if result != 24 {
		t.Error("wrong answer")
	}
}

func TestDivide(t *testing.T) {
	result := Divide(12, 2)
	if result != 6 {
		t.Error("wrong answer")
	}
}

func TestSin(t *testing.T) {
	result := Sin(30)
	if result != math.Sin(30) {
		t.Errorf("wrong answer got %f", result)
	}
}

func TestCos(t *testing.T) {
	result := Cos(30)
	if result != math.Cos(30) {
		t.Errorf("wrong answer got %f", result)
	}
}

func TestTan(t *testing.T) {
	result := Tan(30)
	if result != math.Tan(30) {
		t.Errorf("wrong answer got %f", result)
	}
}

func TestSqrt(t *testing.T) {
	result, _ := Sqrt(30)
	if result != math.Sqrt(30) {
		t.Errorf("wrong answer got %f", result)
	}
}

func TestSqrtOfNegativeNumbers(t *testing.T) {
	_, err := Sqrt(-30)
	if err.Error() != WrongInputError {
		t.Error("should throw error for negative numbers")
	}
}
