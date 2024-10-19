package main

import "testing"

func Sum(a, b int) int

func TestSum(t *testing.T) {
	a := 4
	b := 5
	result := Sum(a, b)

	if result != a+b {
		t.Errorf("expected %v; got %v", a+b, result)
	}
}
