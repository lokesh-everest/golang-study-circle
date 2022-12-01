package calculator

import "testing"

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
