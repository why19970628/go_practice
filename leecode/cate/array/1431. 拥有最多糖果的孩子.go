package main

/*
https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75

给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	resp := make([]bool, len(candies))
	if len(candies) == 0 {
		return resp
	}
	maxVal := candies[0]
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxVal {
			maxVal = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxVal {
			resp[i] = true
		} else {
			resp[i] = false
		}
	}
	return resp
}
