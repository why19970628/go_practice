package mmgp

/*
只能买卖一次

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

*/

/*
暴力
时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func maxProfit_121(prices []int) int {
	resp := 0
	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices); j++ {
			resp = max(resp, prices[j]-prices[i])
		}
	}
	return resp
}

func maxProfit_121_2(prices []int) int {
	resp := 0
	minVal := prices[0]
	for i := 0; i < len(prices); i++ {
		resp = max(resp, prices[i]-minVal)
		minVal = min(minVal, prices[i])
	}
	return resp
}

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-100-liked
*/
