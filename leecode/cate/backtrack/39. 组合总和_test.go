package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	dfscombinationSum(candidates, 0, target)
	return res
}

func dfscombinationSum(candidates []int, start, target int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		dfscombinationSum(candidates, i, target-candidates[i])
		path = path[:len(path)-1]
	}
}

var (
	resp39 [][]int
	temp39 []int
)

func combinationSumV2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resp39 = make([][]int, 0)
	combinationSumV2DFS(candidates, 0, target)
	return resp39
}

func combinationSumV2DFS(candidates []int, start, target int) {
	if target == 0 {
		t := make([]int, len(temp39))
		copy(t, temp39)
		resp39 = append(resp39, t)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[start] > target {
			break
		}
		temp39 = append(temp39, candidates[i])
		combinationSumV2DFS(candidates, i, target-candidates[i])
		temp39 = temp39[:len(temp39)-1]
	}
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSumV2([]int{2, 3, 6, 7}, 7))
}
