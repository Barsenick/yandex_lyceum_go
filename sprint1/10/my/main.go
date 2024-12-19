package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	var bytes [128]byte
	nRead, err1 := r.Read(bytes[:][:n])
	if err1 != nil {
		return err1
	}
	_, err2 := w.Write(bytes[:][:nRead])
	if err2 != nil {
		return err2
	}

	return nil
}
