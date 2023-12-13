package main

/*
931. 下降路径最小和
https://leetcode.cn/problems/minimum-falling-path-sum/description/
给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。
下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列。

[

	[1,2,3],
	[4,5,6],
	[7,8,9]

]
*/

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == m-1 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1])
			} else {
				matrix[i][j] += min(min(matrix[i-1][j], matrix[i-1][j-1]), matrix[i-1][j+1])
			}
		}
	}
	minNum := matrix[n-1][0]
	for i := 0; i < m; i++ {
		if matrix[n-1][i] < minNum {
			minNum = matrix[n-1][i]
		}
	}
	return minNum
}
