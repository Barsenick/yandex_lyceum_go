package main

func SumDigitsRecursive(n int) int {
	if n == 0 {
		return n
	}
	return n%10 + SumDigitsRecursive(n/10)
}
