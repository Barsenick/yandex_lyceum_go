package main

import (
	"errors"
	"io"
)

func WriteString(s string, w io.Writer) error {
	n, err := w.Write([]byte(s))
	if n != len(s) {
		return errors.New("n != len(s)")
	}
	if err != nil {
		return err
	}
	return nil
}
