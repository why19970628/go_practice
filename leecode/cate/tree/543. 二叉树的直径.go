package main

/*
https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
*/
func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lh := dfs(root.Left)
		rh := dfs(root.Right)
		maxDia = max(maxDia, lh+rh)
		return 1 + max(lh, rh)
	}

	dfs(root)
	return maxDia
}
