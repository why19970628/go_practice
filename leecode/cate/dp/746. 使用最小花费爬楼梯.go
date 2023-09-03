package main

import "fmt"

/*
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。


输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15
*/

func minCostClimbingStairs(cost []int) int {
	resp := make([]int, len(cost)+1)
	resp[0] = 0
	resp[1] = 0
	for i := 2; i <= len(cost); i++ {
		resp[i] = min(resp[i-1]+cost[i-1], resp[i-2]+cost[i-2])
	}
	return resp[len(cost)]

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
