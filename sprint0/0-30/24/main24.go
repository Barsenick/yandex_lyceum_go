package main

import "fmt"

//func Clean(nums []int, x int) []int {
//	var cleaned []int = nums
//	for i, val := range cleaned {
//		if val == x {
//			for ii := 0; ii < len(cleaned)-i; ii++ {
//				fmt.Println("cleaned", val)
//				cleaned[i] = cleaned[i+1]
//			}
//		}
//	}
//	return cleaned
//}

func Clean(nums []int, x int) []int {
	ii := 0
	for i, val := range nums {
		if val != x {
			nums[ii] = nums[i]
			ii++
		}
	}
	return nums[:ii]
}

func main() {
	nums := []int{6, 5, 4, 6, 2, 4, 6, 2, 1, 6}
	x := 6
	fmt.Println(Clean(nums, x))
}
