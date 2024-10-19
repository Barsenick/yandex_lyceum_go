package main

import "fmt"

func FindMaxKey(m map[int]int) int {
	maxKey := 0
	haveFoundSomethingYet := false

	for i := range m {
		if !haveFoundSomethingYet {
			maxKey = i
			haveFoundSomethingYet = true
			continue
		}
		if i > maxKey {
			maxKey = i
		}
	}
	return maxKey
}

func main() {
	m := map[int]int{
		10: 37,
		3:  19,
	}

	fmt.Printf("Max. key - %d", FindMaxKey(m))
}
