package main

import "fmt"

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

// 有序数组！！！

// 双指针法
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	k := len(nums) - 1
	for left <= right {
		l := nums[left] * nums[left]
		r := nums[right] * nums[right]
		if r > l {
			ans[k] = r
			right--
		} else {
			ans[k] = l
			left++
		}
		k--
	}

	return ans
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
