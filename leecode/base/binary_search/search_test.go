package main

import (
	"fmt"
	"testing"
)

// 二分查找
func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			mid = left + 1
		}
		if nums[mid] > target {
			mid = right - 1
		}
	}
	return -1
}

func TestBinary_search(t *testing.T) {
	fmt.Println(BinarySearch([]int{2, 3, 6, 7}, 3))
}
