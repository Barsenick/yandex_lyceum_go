package main

import (
	"bytes"
	"fmt"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	var bytes [128]byte
	_, err := r.Read(bytes[:])
	if err != nil {
		return false, err
	}
	seqFound := 0
	for i := 0; i < len(bytes); i++ {
		if seqFound < len(seq) && bytes[i] == seq[seqFound] {
			seqFound++
			continue
		}
		if seqFound == len(seq) {
			return true, nil
		} else {
			seqFound = 0
		}
	}
	if seqFound == len(bytes)-1 {
		return true, nil
	}
	return false, err
}

func main() {
	reader := bytes.NewReader([]byte("Hello, World!"))
	contains, err := Contains(reader, []byte("Hello, World!"))
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %v", contains)
	}
}
