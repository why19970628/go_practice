package backtrack

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/subsets/description/

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]

*/

//func subsets(nums []int) [][]int {
//	res, path = make([][]int, 0), make([]int, 0)
//	res = append(res, []int{})
//	subsetsDfs(nums, 0)
//	return res
//}
//
//func subsetsDfs(nums []int, start int) {
//	if len(path) > 0 {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//	}
//
//	for i := start; i < len(nums); i++ {
//		path = append(path, nums[i])
//		subsetsDfs(nums, i+1)
//		path = path[:len(path)-1]
//	}
//}

func subsets(nums []int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	res = append(res, []int{})
	subsetsDfs(nums, 0)
	return res
}

func subsetsDfs(nums []int, start int) {
	if len(path) > 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		subsetsDfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

var (
	subsetsV2Resp [][]int
	subsetsV2Temp []int
)

func subsetsV2(nums []int) [][]int {
	subsetsV2Resp = make([][]int, 0)
	subsetsV2Temp = make([]int, 0)
	subsetsV2Dfs(nums, 0)
	subsetsV2Resp = append(subsetsV2Resp, []int{})
	return subsetsV2Resp
}

func subsetsV2Dfs(nums []int, start int) {
	if len(subsetsV2Temp) > 0 {
		t := make([]int, len(subsetsV2Temp))
		copy(t, subsetsV2Temp)
		subsetsV2Resp = append(subsetsV2Resp, t)
	}
	for i := start; i < len(nums); i++ {
		subsetsV2Temp = append(subsetsV2Temp, nums[i])
		subsetsV2Dfs(nums, i+1)
		subsetsV2Temp = subsetsV2Temp[:len(subsetsV2Temp)-1]
	}
	return
}

func TestSubsets(t *testing.T) {
	fmt.Println(subsetsV2([]int{1, 2, 3}))
}
