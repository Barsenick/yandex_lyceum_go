package main

func isLatin(input string) bool {
	for _, r := range input {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
