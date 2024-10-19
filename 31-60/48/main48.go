package main

import (
	"errors"
)

var (
	ErrNegativeNumberFactorial = errors.New("factorial is not defined for negative numbers")
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeNumberFactorial
	}

	if n == 0 {
		return 1, nil
	}

	f := 1

	for i := 1; i <= n; i++ {
		f *= i
	}

	return f, nil
}
