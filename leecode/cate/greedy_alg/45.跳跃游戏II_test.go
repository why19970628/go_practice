package greedy_alg

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	var count = 0
	next := 0
	cur := 0

	for i := 0; i < n-1; i++ {
		next = max(i+nums[i], next)
		if i == cur {
			cur = next
			count++
		}
	}

	return count
}
