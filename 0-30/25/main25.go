package main

func SliceCopy(nums []int) []int {
	var newSlice []int = make([]int, len(nums))
	copy(newSlice, nums)
	return newSlice
}
