package main

import "fmt"

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// 典型的动态规划题目
//确定初始条件: f(1) = 1, f(2) = 2
//确定状态方程: f(n) = f(n-1) + f(n-2)
//确定最终答案: f(n) 即为最终的答案
//
// https://leetcode-cn.com/problems/climbing-stairs/solution/leetcode_0070-by-ssezhangpeng-tucy/
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	before, after := 1, 2
	for i := 3; i <= n; i++ {
		before, after = after, before+after
		fmt.Println(i, before, after)
	}

	return after
}

func main() {
	fmt.Println(climbStairs2(10))
}
