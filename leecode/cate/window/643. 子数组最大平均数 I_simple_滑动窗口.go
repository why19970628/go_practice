package main

/*
https://leetcode.cn/problems/maximum-average-subarray-i/description/?envType=study-plan-v2&envId=leetcode-75

给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。



示例 1：

输入：nums = [1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
示例 2：

输入：nums = [5], k = 1
输出：5.00000
*/

func findMaxAverage(nums []int, k int) float64 {
	maxVal, sum := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxVal = sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > maxVal {
			maxVal = sum
		}
	}
	return float64(maxVal) / float64(k)
}
