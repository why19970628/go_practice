package main

import "fmt"

/*
https://leetcode.cn/problems/next-greater-element-i/

nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	mp := map[int]int{}
	for i, v := range nums1 {
		mp[v] = i
	}
	// 单调栈
	stack := []int{0}

	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			// 要移除的的栈顶index
			top := stack[len(stack)-1]
			if _, ok := mp[nums2[top]]; ok { // 看map里是否存在这个元素
				index := mp[nums2[top]] // 根据map找到nums2[top] 在 nums1中的下标
				res[index] = nums2[i]
			}

			// 比当前数据小的栈顶 全部出栈
			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, i)
	}
	return res
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	resp := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		resp[i] = -1
	}
	mp := make(map[int]int)
	for index, v := range nums1 {
		mp[v] = index
	}
	stack := make([]int, 0)
	for index2, v2 := range nums2 {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < v2 {
			topIndex := stack[len(stack)-1]
			topVal := nums2[topIndex]
			index1, ok := mp[topVal]
			if ok {
				resp[index1] = v2
			}
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index2)
	}
	return resp
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	// [-1,3,-1]
	fmt.Println(nextGreaterElement2(nums1, nums2))
}
