package main

import "fmt"

/*
https://leetcode.cn/problems/kth-largest-element-in-an-array/

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

*/
func findKthLargest(nums []int, k int) int {
	return quickSortV2(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSortV2(nums []int, l, r, target int) int {
	pIndex := partition2(nums, l, r)
	fmt.Println("pIndex:", pIndex)
	if target == pIndex {
		return nums[target]
	}
	if target > pIndex {
		return quickSortV2(nums, pIndex+1, r, target)
	} else {
		return quickSortV2(nums, l, pIndex-1, target)
	}
}

func partition2(nums []int, l, r int) int {
	p := nums[r]
	i, j := l, l
	for j < len(nums) {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
