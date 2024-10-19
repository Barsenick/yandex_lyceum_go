package main

import (
	"errors"
	"fmt"
)

var (
	ErrPosOutOfRange = errors.New("position out of range")
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	fmt.Println(len(str))
	if position < 0 || position >= len([]rune(str)) {
		return rune(0), ErrPosOutOfRange
	}
	return []rune(str)[position], nil
}
