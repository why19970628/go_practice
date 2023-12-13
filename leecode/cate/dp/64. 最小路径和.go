package main

import "fmt"

/*
https://leetcode.cn/problems/minimum-path-sum/

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
*/
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				grid[i][j] += minV2(grid[i-1][j], grid[i][j-1])
			}

		}
	}
	return grid[m-1][n-1]
}

func minPathSumV2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += minV2(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func minV2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	a := [][]int{
		{
			1, 3, 1,
		},
		{
			1, 5, 1,
		}, {
			4, 2, 1,
		},
	}
	fmt.Println(minPathSum(a))
}
