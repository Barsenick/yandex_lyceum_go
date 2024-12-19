package main

import "errors"

var (
	ErrDivisionByZero = errors.New("division by zero is not allowed")
)

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return float64(a / b), nil
}
