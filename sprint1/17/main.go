package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err1 := os.Open(inputFilename)
	if err1 != nil {
		return err1
	}
	defer inputFile.Close()
	_, err2 := inputFile.Seek(int64(startpos), 0)
	if err2 != nil {
		return err2
	}
	data := make([]byte, 0)
	buffer := make([]byte, 256)
	for {
		n, err3 := inputFile.Read(buffer)
		if err3 != nil {
			if err3 == io.EOF {
				break
			}
			return err3
		}
		data = append(data, buffer[:n]...)
		_, err4 := inputFile.Seek(256, 0)
		if err4 != nil {
			return err4
		}
	}

	err5 := os.WriteFile(outFileName, data, 0644)
	if err5 != nil {
		return err5
	}
	return nil
}
