package main

func IsPalindrome(input string) bool {
	rs := []rune(input)
	new := []rune{}

	for i := 0; i < len(rs); i++ {
		if rs[i] != ' ' {
			new = append(new, rs[i])
		}
	}

	for i := 0; i < len(new)/2; i++ {
		if new[i] != new[len(new)-1-i] {
			return false
		}
	}
	return true
}
