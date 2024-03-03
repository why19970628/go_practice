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

func threeSumV3(nums []int) [][]int {

	resp := make([][]int, 0)
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	mp := make(map[int]bool)
	// -4,-1,-1, 0,1,2
	for i := 0; i < len(nums); i++ {
		// 未去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		mp[nums[i]] = true
		ans := twoSumV3(nums, i+1, 0-nums[i])
		if len(ans) > 0 {
			for _, an := range ans {
				resp = append(resp, []int{nums[i], nums[an[0]], nums[an[1]]})
			}
		}
	}
	return resp
}

// -4,-1,-1, 0,1,2
func twoSumV3(nums []int, start, target int) [][]int {
	ans := make([][]int, 0)
	for i := start; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 去重
			if nums[start] == nums[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -4, 8}))
}
