package main

import "fmt"

/*
https://leetcode.cn/problems/median-of-two-sorted-arrays/submissions/

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 1. 合并数组
	// 2. 排序
	// 3. 求中位数
	nums := mergeSortedArrays(nums1, nums2)
	fmt.Println("nums", nums)
	return getMedian(nums)
}

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	newArr := []int{}
	count := len(nums1) + len(nums2)
	for count > 0 {
		if len(nums1) == 0 {
			newArr = append(newArr, nums2...)
			break
		}
		if len(nums2) == 0 {
			newArr = append(newArr, nums1...)
			break
		}
		if nums1[0] < nums2[0] {
			newArr = append(newArr, nums1[0])
			nums1 = nums1[1:]
		} else {
			newArr = append(newArr, nums2[0])
			nums2 = nums2[1:]
		}
		count--
	}
	return newArr
}

func getMedian(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2-1]) + float64(nums[len(nums)/2])) / float64(2)
	}
	return float64(nums[len(nums)/2])
}
