package main

import (
	"sort"
)

func Permutations(input string) []string {
	result := []string{}

	if len(input) == 0 {
		return result
	}
	if len(input) == 1 {
		return []string{input}
	}

	for i, c := range input {
		for _, perm := range Permutations(input[:i] + input[i+1:]) {
			result = append(result, string(c)+perm)
		}
	}

	sort.Strings(result)
	return result
}
