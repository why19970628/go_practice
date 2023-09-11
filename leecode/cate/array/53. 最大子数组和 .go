package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23
*/

func getMaxSum(arr []int) int {
	var sum, maxSum int
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func maxSubArrayV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	resp := -math.MaxInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		resp = max(sum, resp)
		if sum < 0 {
			sum = 0
		}
	}
	return resp

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := -math.MaxInt32
	var count = 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}
	return result
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(getMaxSum(nums))
}
