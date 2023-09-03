package greedy_alg

import "math"

/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:  连续子数组  [4,-1,2,1] 的和最大，为  6。

*/

func maxSubArrayV3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := -math.MaxInt32
	var count = 0
	for i := 1; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		// 局部最优的情况下，并记录最大的"连续和",可以推出全局最优。
		//从代码角度上来讲：遍历 nums，从头开始用 count 累积，如果 count 一旦加上 nums[i]变为负数，那么就应该从 nums[i+1]开始从 0 累积 count 了，因为已经变为负数的 count，只会拖累总和。
		if count < 0 {
			count = 0
		}
	}

	return result
}

func maxSubArrayV2(nums []int) int {
	resp := nums[0]
	for i := 1; i < len(nums); i++ {
		//当i元素+i-1元素大于当前i元素就加（决定了nums[i-1]>0）
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > resp {
			resp = nums[i]
		}
	}
	return resp
}

// 暴力解法的思路，第一层 for 就是设置起始位置，第二层 for 循环遍历数组寻找最大值
func maxSubArray(nums []int) int {
	resp := -math.MaxInt32
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i + 1; j < len(nums); j++ {
			temp += nums[j]
			if temp > resp {
				resp = temp
			}
		}
	}

	return resp
}
