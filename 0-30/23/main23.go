package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil {
		return nil, errors.New("nums пуст")
	}
	var slice []int = make([]int, 0, n)

	i := 0
	for _, val := range nums {
		if val < limit {
			slice = append(slice, val)
			i++
		}
		if i == n {
			break
		}
	}

	return slice, nil

}

func main() {
	var nums []int = []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	var n int = 5
	var limit int = 3

	fmt.Println(UnderLimit(nums, limit, n))
}
