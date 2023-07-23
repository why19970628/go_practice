package backtrack

import (
	"fmt"
	"testing"
)

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]

*/

func subsets(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
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
func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}
