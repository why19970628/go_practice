package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] == target {
			return mid
		}

	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{-1, 0, 1, 2, 3, 9}, 9))
}
