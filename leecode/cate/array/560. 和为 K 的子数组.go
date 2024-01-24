package main

/*
https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2
*/

// 前缀和
func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	resp := 0
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if mp[preSum-k] > 0 {
			resp += mp[preSum-k]
		}
		mp[preSum]++
	}
	return resp
}
