package main

/*

https://leetcode.cn/problems/climbing-stairs/
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

*/

func climbStairs(n int) int {
	bp := make([]int, n+1)
	bp[1] = 1
	bp[2] = 2
	for i := 3; i <= n; i++ {
		bp[i] = bp[i-1] + bp[i-2]
	}
	return bp[n]
}

func climbStairsV2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
