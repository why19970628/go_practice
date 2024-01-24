package main

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]

*/

func generate(numRows int) [][]int {
	resp := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		lastRow := resp[len(resp)-1]
		next := make([]int, i+1)
		next[0], next[len(next)-1] = 1, 1
		for j := 1; j < i; j++ {
			next[j] = lastRow[j-1] + lastRow[j]
		}
		resp = append(resp, next)
	}
	return resp
}
