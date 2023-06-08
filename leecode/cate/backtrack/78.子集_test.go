package backtrack

import (
	"fmt"
	"testing"
)

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
