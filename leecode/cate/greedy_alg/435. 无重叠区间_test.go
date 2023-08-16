package greedy_alg

import "sort"

/*
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
*/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			// 不重叠个数
			res++
		} else {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}

	}
	return len(intervals) - res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
