package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, index}
		} else {
			m[v] = index
		}
	}
	return []int{}
}

func twoSumV2(nums []int, target int) []int {
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := mp[target-nums[i]]; ok {
			return []int{v, i}
		}
		mp[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
