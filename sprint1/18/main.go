package main

import "os"

func ModifyFile(filename string, pos int, val string) {
	f, err1 := os.OpenFile(filename, os.O_WRONLY, 0777)
	if err1 != nil {
		return
	}
	defer f.Close()
	f.Seek(int64(pos), 0)
	_, err2 := f.Write([]byte(val))
	if err2 != nil {
		return
	}
}
