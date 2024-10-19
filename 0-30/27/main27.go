package main

func Mix(nums []int) []int {
	n := len(nums) / 2
	mixed := make([]int, n*2)

	for i := 0; i < n; i++ {
		mixed[i*2] = nums[i]
		mixed[i*2+1] = nums[n+i]
	}

	return mixed
}
