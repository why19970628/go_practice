package main

import "fmt"

/*
https://leetcode.cn/problems/longest-increasing-subsequence/description/

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1
*/
func lengthOfLIS_DP(nums []int) int {
	dp := make([]int, len(nums))

	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}

	resp := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		resp = max(resp, dp[i])
	}
	return resp
}

/*
时间复杂度：O(n^2),其中 n为nums的长度。
空间复杂度：O(n)。
*/
func lengthOfLIS(nums []int) (ans int) {
	f := make([]int, len(nums))
	for i, x := range nums {
		// 内部循环中的 j 遍历 i 之前的元素，y 是遍历到的元素值。
		// 在这个内部循环中，检查是否有比当前元素 x 小的元素 y。
		for j, y := range nums[:i] {
			// 如果存在比 x 小的元素 y，那么更新 f[i]，使其成为 f[j] 和当前 f[i] 中的最大值。
			// 这是因为在 y < x 的情况下，x 可以加入到以 y 结尾的子序列中，形成更长的递增子序列。
			if y < x {
				f[i] = max(f[i], f[j])
			}
		}
		// 每当找到一个满足条件的 y，f[i] 就会增加1，表示以当前元素 x 结尾的最长递增子序列的长度。
		f[i]++
		ans = max(ans, f[i])
	}
	return
}

func main() {
	nums := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}
