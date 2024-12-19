package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	i := 0
	if lineNum < 0 {
		return ""
	}
	for fileScanner.Scan() {
		if i == lineNum {
			return string(fileScanner.Text())
		}
	}
	return ""
}
