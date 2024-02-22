package main

import "fmt"

/*
给一个数组，找出一个连续子数组，要求子数组中的每个数字只出现一次，输出这个子数组的长度最长是多少。
*/

func maxUniqueSubarrayLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength := 0
	seen := make(map[int]bool)
	left := 0

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] {
			delete(seen, nums[left])
			left++
		}
		seen[nums[right]] = true
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func main() {
	nums := []int{1, 2, 3, 4, 1, 2, 3, 5}
	fmt.Println(maxUniqueSubarrayLength(nums)) // 输出为 5，对应的子数组是 [1, 2, 3, 4, 5]
}
