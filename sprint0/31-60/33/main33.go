package main

func SumOfValuesInMap(m map[int]int) int {
	sum := 0
	for i := range m {
		sum += m[i]
	}

	return sum
}
