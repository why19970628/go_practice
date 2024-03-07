package main

/*
https://leetcode.cn/problems/max-consecutive-ones-iii/description/?envType=study-plan-v2&envId=leetcode-75

给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。

示例 1：

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

*/

// [1,1,1,0,0,0,1,1,1,1,0]

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	zeroCount := 0

	maxVal := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}

		// 移动左指针
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		right++
		if right-left > maxVal {
			maxVal = right - left
		}
	}
	return maxVal
}

func main() {

}
