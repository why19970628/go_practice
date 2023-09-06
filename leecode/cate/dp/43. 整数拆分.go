package main

import (
	"fmt"
)

/*
给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
返回 你可以获得的最大乘积 。

示例 1:

输入: n = 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: n = 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

https://leetcode.cn/problems/integer-break/
*/

func integerBreak(n int) int {
	/**
	  动态五部曲
	  1.确定dp下标及其含义
	  2.确定递推公式
	  3.确定dp初始化
	  4.确定遍历顺序
	  5.打印dp
	  **/
	// 每个i最大面积
	dp := make([]int, n+1)
	// n=1 || n=2时，面积为1
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
		//fmt.Println(dp)
	}
	return dp[n]
}

func integerBreakV2(n int) int {
	if n <= 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; i < i; j++ {
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(integerBreak(10))
}
