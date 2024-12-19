package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	strStart := start.Format("02.01.2006")
	strEnd := end.Format("02.01.2006")

	logs := make([]string, 0)

	f, err := os.Open(inputFileName)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	i := 0

	foundFirstDate := false

	for fileScanner.Scan() {
		if !foundFirstDate {
			date, log, found := strings.Cut(fileScanner.Text(), " ")
			if !found {
				return []string{}, errors.New("wrong format")
			}
			if date == strStart {
				foundFirstDate = true
				logs = append(logs, date+" "+log)
				continue
			}
		} else {
			date, log, found := strings.Cut(fileScanner.Text(), " ")
			if !found {
				return []string{}, errors.New("wrong format")
			}
			logs = append(logs, date+" "+log)
			if date == strEnd {
				break
			}
		}
		i++
	}

	if !foundFirstDate {
		return []string{}, errors.New("didnt find first date")
	}

	if len(logs) == 0 {
		return []string{}, errors.New("empty logs")
	}

	return logs, nil
}
