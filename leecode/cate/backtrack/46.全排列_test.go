package backtrack

import (
	"fmt"
	"testing"
)

var st []bool // state的缩写

func permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	permuteDfs(nums, 0)
	return res
}

func permuteDfs(nums []int, cur int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	// 这里不能用cur当起点，因为可以有重复数据
	for i := 0; i < len(nums); i++ {
		if !st[i] {
			st[i] = true
			path = append(path, nums[i])
			permuteDfs(nums, cur+1)
			path = path[:len(path)-1]
			st[i] = false
		}

	}
}

var (
	resp46 [][]int
	temp46 []int
	mp2    = make(map[int]bool)
)

func permuteV2(nums []int) [][]int {
	resp46 = make([][]int, 0)
	mp2 = make(map[int]bool, len(nums))
	permuteV2DFS(nums, 0)
	return resp46
}

func permuteV2DFS(nums []int, start int) {
	if len(temp46) == len(nums) {
		t := make([]int, len(nums))
		copy(t, temp46)
		resp46 = append(resp46, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		//if i == start {
		//	continue
		//}

		if val, ok := mp2[nums[i]]; !ok || val == false {
			mp2[nums[i]] = true
			temp46 = append(temp46, nums[i])
			permuteV2DFS(nums, i)
			temp46 = temp46[:len(temp46)-1]
			mp2[nums[i]] = false
		}
	}
}

func TestPermute(t *testing.T) {
	fmt.Println(permuteV2([]int{1, 2, 3}))
}
