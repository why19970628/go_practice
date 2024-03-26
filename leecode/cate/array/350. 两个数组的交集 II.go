package main

import "fmt"

/*
https://leetcode.cn/problems/intersection-of-two-arrays-ii?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

*/

func intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for _, v := range nums1 {
		mp[v]++
	}
	resp := make([]int, 0)
	for _, v := range nums2 {
		val, ok := mp[v]
		if ok {
			if val >= 1 {
				resp = append(resp, v)
				mp[v]--
			}
		}
	}
	return resp
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
