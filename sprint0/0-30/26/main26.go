package main

func Join(nums1, nums2 []int) []int {
	var joinedSlice []int = make([]int, len(nums1)+len(nums2))
	copy(joinedSlice, nums1)
	for i, val := range nums1 {
		joinedSlice[len(nums1)+i] = val
	}
	return joinedSlice
}
