package main

import "fmt"

/*
https://leetcode.cn/problems/next-permutation/

整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。
*/

func nextPermutation(nums []int) {
	for i := 0; i < len(nums); i++ {
		index := 0
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] > nums[i] && nums[j] < nums[j+1] {
				index = j
			}
		}
		if index > i && index != i+1 {
			nums[i+1], nums[index] = nums[index], nums[i+1]
		}

	}
	return
}

/*

https://leetcode.cn/problems/next-permutation/solutions/479324/jie-fa-hen-jian-dan-jie-shi-qi-lai-zen-yao-jiu-na-/?envType=study-plan-v2&envId=top-100-liked

思路
如何变大：从低位挑一个大一点的数，交换前面一个小一点的数。
变大的幅度要尽量小。
像 [3,2,1] 这样递减的，没有下一个排列，已经稳定了，没法变大。
像 [1,5,2,4,3,2] 这种，怎么稍微变大？

从右往左，寻找第一个比右邻居小的数，把它换到后面去

“第一个”意味着尽量是低位，“比右邻居小”意味着它是从右往左的第一个波谷
比如，1 5 (2) 4 3 2，中间这个 2。

接着还是从右往左，寻找第一个比这个 2 大的数。15 (2) 4 (3) 2，交换后：15 (3) 4 (2) 2。
还没结束！变大的幅度可以再小一点，仟位微变大了，后三位可以再小一点。
后三位肯定是递减的，翻转，变成[1,5,3,2,2,4]，即为所求。

*/

func nextPermutationV2(nums []int) {
	i := len(nums) - 2
	// 从右往左找到第一个波谷（左边比右边小）
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 从右往左，找到第一个比i大点数字，交换
	j := len(nums) - 1

	if i >= 0 {
		for j > 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	l, r := i+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	arr := []int{1, 2, 3} // [1,3,2]
	nextPermutation(arr)
	fmt.Println(arr)
}
