package main

func FiveSteps(array [5]int) [5]int {
	var newArray [5]int
	for i := len(array); i >= 0; i-- {
		newArray[5-i] = array[i]
	}
	return newArray
}
