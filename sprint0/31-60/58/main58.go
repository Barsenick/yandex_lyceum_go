package main

import (
	"testing"
)

func Multiply(a, b int) int

func TestMultiply(t *testing.T) {
	a := 4
	b := 5
	result := Multiply(a, b)

	if result != a+b {
		t.Errorf("expected %v; got %v", a*b, result)
	}
}
