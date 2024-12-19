package main

func CalculateSeriesSum(n int) float64 {
	var sum float64 = 0

	for i := 1; i <= n; i++ {
		sum += float64(1) / float64(i)
	}

	return sum
}
