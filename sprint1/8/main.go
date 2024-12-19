package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024) // Allocate a buffer with a reasonable size
	n, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(data[:n]), nil
}
