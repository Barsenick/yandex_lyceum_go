package main

import (
	"errors"
	"fmt"
)

var (
	ErrNegNumber = errors.New("negative numbers are not allowed")
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", ErrNegNumber
	}

	return fmt.Sprintf("%b", num), nil
}
