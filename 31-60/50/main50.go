package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrNotInts = errors.New("invalid input, please provide two integers")
)

func SumTwoIntegers(a, b string) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	inta, err := strconv.Atoi(a)
	if err != nil {
		return 0, ErrNotInts
	}

	intb, err := strconv.Atoi(b)
	if err != nil {
		return 0, ErrNotInts
	}

	return inta + intb, nil
}
