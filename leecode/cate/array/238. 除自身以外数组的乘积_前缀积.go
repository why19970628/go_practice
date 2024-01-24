package main

/*

https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-100-liked

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
*/

// 前缀积： 维护ans[i]: 为左侧所有元素之积 * 右侧所有元素之积
func productExceptSelf(nums []int) []int {
	resp := make([]int, 0)
	resp[0] = 1
	for i := 1; i < len(nums); i++ {
		resp[i] = resp[i-1] * nums[i-1]
	}
	sum := 1
	for j := len(nums) - 1; j >= 0; j-- {
		resp[j] = resp[j] * sum
		sum = sum * nums[j]
	}
	return resp
}
