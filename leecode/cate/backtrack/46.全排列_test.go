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
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
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

//func permuteDfs(nums []int, cur int) {
//	if cur == len(nums) {
//		tmp := make([]int, len(path))
//		copy(tmp, path)
//		res = append(res, tmp)
//	}
//	for i := 0; i < len(nums); i++ {
//		if !st[i] {
//			path = append(path, nums[i])
//			st[i] = true
//			permuteDfs(nums, cur+1)
//			st[i] = false
//			path = path[:len(path)-1]
//		}
//	}
//}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{4, 6, 7}))
}
