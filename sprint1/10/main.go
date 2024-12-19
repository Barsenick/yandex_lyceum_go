package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	var buf [64]byte
	for i := 0; uint(i) < n; i += len(buf) {
		nRead, err := r.Read(buf[:])
		if err != nil && err != io.EOF {
			return err
		}
		if nRead == 0 {
			break
		}
		if _, err := w.Write(buf[:nRead]); err != nil {
			return err
		}
	}
	return nil
}
