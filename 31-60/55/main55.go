package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {

	switch start.Day() {
	case int(time.Monday):
		return start.AddDate(0, 0, 2)
	case int(time.Tuesday):
		return start.AddDate(0, 0, 2)

	case int(time.Wednesday):
		return start.AddDate(0, 0, 2)

	case int(time.Thursday):
		return start.AddDate(0, 0, 2)

	case int(time.Friday):
		return start.AddDate(0, 0, 4)

	case int(time.Saturday):
		return start.AddDate(0, 0, 3)

	case int(time.Sunday):
		return start.AddDate(0, 0, 2)
	default:
		return start.AddDate(0, 0, 2)
	}
}

func main() {
	start := time.Date(2023, time.October, 6, 0, 0, 0, 0, time.UTC) // Friday
	fmt.Printf("Start: %v\n", start.Weekday())
	nextWorkday := NextWorkday(start)
	fmt.Printf("Result: %v\n", nextWorkday.Weekday())
	expected := time.Date(2023, time.October, 9, 0, 0, 0, 0, time.UTC) // The following Monday
	fmt.Printf("Expected: %v\n", expected.Weekday())
	if !nextWorkday.Equal(expected) {
		fmt.Printf("Expected next workday to be %v, but got %v", expected, nextWorkday)
	}
}
