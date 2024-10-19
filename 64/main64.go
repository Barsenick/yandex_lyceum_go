package main

func SumDigitsRecursive(n int) int {
	if n == 0 {
		return n
	}
	return n%10 + SumDigitsRecursive(n/10)
}

func CalculateDigitalRoot(n int) int {
	num := 0
	for {
		if num <= 9 {
			return num
		}
		num += SumDigitsRecursive(n)
	}
}
