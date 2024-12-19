package main

import (
	"errors"
	"io"
	"strings"
)

type UpperWriter struct {
	UpperString string
	io.Writer
}

func (uw *UpperWriter) Write(b []byte) (int, error) {
	if strings.ToValidUTF8(string(b), "") != string(b) {
		return 0, errors.New("invalid string")
	}
	uw.UpperString = strings.ToUpper(string(b))
	return len(string(b)), nil
}
