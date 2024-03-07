package main

/*

https://leetcode.cn/problems/find-the-difference-of-two-arrays/description/?envType=study-plan-v2&envId=leetcode-75

给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：

answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
注意：列表中的整数可以按 任意 顺序返回。



示例 1：

输入：nums1 = [1,2,3], nums2 = [2,4,6]
输出：[[1,3],[4,6]]
解释：
对于 nums1 ，nums1[1] = 2 出现在 nums2 中下标 0 处，然而 nums1[0] = 1 和 nums1[2] = 3 没有出现在 nums2 中。因此，answer[0] = [1,3]。
对于 nums2 ，nums2[0] = 2 出现在 nums1 中下标 1 处，然而 nums2[1] = 4 和 nums2[2] = 6 没有出现在 nums2 中。因此，answer[1] = [4,6]。
示例 2：

输入：nums1 = [1,2,3,3], nums2 = [1,1,2,2]
输出：[[3],[]]
解释：
对于 nums1 ，nums1[2] 和 nums1[3] 没有出现在 nums2 中。由于 nums1[2] == nums1[3] ，二者的值只需要在 answer[0] 中出现一次，故 answer[0] = [3]。
nums2 中的每个整数都在 nums1 中出现，因此，answer[1] = [] 。

*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	mp1 := make(map[int]bool)
	mp2 := make(map[int]bool)
	resp := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		mp1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		mp2[nums2[i]] = true
	}

	a1 := []int{}
	for key, _ := range mp1 {
		if !mp2[key] {
			a1 = append(a1, key)
		}
	}
	a2 := []int{}
	for key, _ := range mp2 {
		if !mp1[key] {
			a2 = append(a2, key)
		}
	}
	resp = append(resp, a1)
	resp = append(resp, a2)
	return resp
}
