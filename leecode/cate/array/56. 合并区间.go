package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-100-liked

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func mergeArray(intervals [][]int) [][]int {
	resp := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			resp = append(resp, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	resp = append(resp, []int{left, right}) // 将最后一个区间放入
	return resp
}

func main() {
	//n1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//n2 := [][]int{{1, 4}, {4, 5}}
	n3 := [][]int{{1, 4}, {2, 3}}

	resp := mergeArray(n3)
	fmt.Println(resp)
}
