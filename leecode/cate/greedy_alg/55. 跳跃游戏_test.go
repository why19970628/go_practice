package greedy_alg

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/jump-game/description/
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

*/

// 贪心算法局部最优解：每次取最大跳跃步数（取最大覆盖范围），整体最优解：最后得到整体最大覆盖范围，看是否能到终点。
// 局部最优推出全局最优，找不出反例，试试贪心！
func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover > n {
			return true
		}
	}
	return false
}

// https://leetcode.cn/problems/jump-game/solutions/465202/55-tiao-yue-you-xi-tan-xin-jing-dian-ti-mu-xiang-j/
func canJumpV2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	cover := 0
	for i := 0; i < cover; i++ {
		cover = max(cover, i+nums[i])
		if cover > len(nums) {
			return true
		}
	}
	return false
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4, 2, 3, 1, 1, 4}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
