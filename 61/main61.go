package main

import (
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	s1 := strings.ToLower(str1)
	s2 := strings.ToLower(str2)

	for _, r := range []rune(s1) {
		if strings.ContainsRune(s2, r) {
			return false
		}
	}
	return true
}
