package main

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	resp := 1
	for k, _ := range mp {
		// 若该数前面不存在数，则把该数当作连续序列的第一个数
		if !mp[k-1] {
			long := 1
			for mp[k+1] {
				long++
				k++
			}
			resp = max(resp, long)
		}
	}
	return resp
}
