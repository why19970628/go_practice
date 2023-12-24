package main

/*
https://leetcode.cn/problems/house-robber-ii/description/

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [1,2,3]
输出：3


如果偷 nums[0]，那么 nums[1] 和 nums[n−1] 不能偷，问题变成从 nums[2]到 nums[n−2] 的非环形版本，调用 198 题的代码解决；
如果不偷 nums[0]，那么问题变成从 nums[1] 到 nums[n−1] 的非环形版本，同样调用 198 题的代码解决。
*/

// 198. 打家劫舍的逻辑
func rob1(nums []int, start, end int) int { // [start,end) 左闭右开
	f0, f1 := 0, 0
	for i := start; i < end; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}
func rob2_v1(nums []int) int {
	n := len(nums)
	return max(nums[0]+rob1(nums, 2, n-1), rob1(nums, 1, n))
}

// 打家劫舍Ⅱ 动态规划
// 时间复杂度O(n) 空间复杂度O(n)
func rob2_v2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result1 := robRange(nums, 0)
	result2 := robRange(nums, 1)
	return max(result1, result2)
}

// 偷盗指定的范围
func robRange(nums []int, start int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[start]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1+start], dp[i-1])
	}
	return dp[len(nums)-1]
}

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	// 创建两个动态规划数组，分别表示不偷第一个房屋和不偷最后一个房屋的情况下的最大金额
	dp1 := make([]int, n) // 不偷第一个房屋
	dp2 := make([]int, n) // 不偷最后一个房屋

	// 不偷第一个房屋，偷1~n 初始化第一个值
	dp1[0] = 0
	dp1[1] = nums[1]

	// 不偷最后一个房屋,偷0~n-1 初始化第一个值
	dp2[0] = nums[0]
	dp2[1] = max(nums[0], nums[1])

	// 计算不偷第一个房屋的情况下的最大金额
	for i := 2; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}

	// 计算不偷最后一个房屋的情况下的最大金额
	for i := 2; i < n-1; i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}

	// 返回两种情况下的最大金额中的较大值
	return max(dp1[n-1], dp2[n-2])
}
