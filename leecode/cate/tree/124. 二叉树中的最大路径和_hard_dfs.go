package main

import "math"

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt64
	maxPathSum_dfs(root)
	return maxSum
}

func maxPathSum_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxPathSum_dfs(root.Left)
	right := maxPathSum_dfs(root.Right)
	maxSum = max2(maxSum, root.Val+left+right)
	ret := root.Val + max2(left, right)
	return max2(ret, 0)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
