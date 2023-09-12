package main

import (
	"fmt"
	"sort"
	"testing"
)

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//[-4,-1,-1,0,1,2]
	resp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := twoSumV2(nums, i+1, 0-nums[i])
		if len(r) > 0 {
			for _, v := range r {
				resp = append(resp, []int{nums[i], v[0], v[1]})
			}
		}
	}
	return resp
}

func twoSumV2(numbers []int, start int, target int) [][]int {
	j := len(numbers) - 1
	resp := make([][]int, 0)
	for i := start; i < len(numbers); i++ {
		if i > start && numbers[i] == numbers[i-1] {
			continue
		}
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			resp = append(resp, []int{numbers[i], numbers[j]})
		}
	}
	return resp
}
