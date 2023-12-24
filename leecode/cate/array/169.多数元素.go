package main

import "sort"

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
输入：nums = [3,2,3]
输出：3
示例 2：
输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

// 先排序，然后返回排好序数组中间的那个数即可
func majorityElement(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return nums[n/2]
}

// map
func majorityElementV2(nums []int) int {
	mp := make(map[int]int, 0)
	for _, num := range nums {
		mp[num]++
		if mp[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

/*
算法描述
摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。
算法可以分为两个阶段：
对抗阶段：分属两个候选人的票数进行两两对抗抵消
计数阶段：计算对抗结果中最后留下的候选人票数是否有效
*/
func majorityElementV3(nums []int) int {
	major, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count++
		} else {
			count--
		}
	}
	return major
}
