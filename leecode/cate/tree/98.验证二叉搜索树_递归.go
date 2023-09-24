package main

import "math"

/*
https://leetcode.cn/problems/validate-binary-search-tree/
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return isValidBSTSolution(root, math.MaxInt64, math.MinInt64)
}

func isValidBSTSolution(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTSolution(root.Left, root.Val, min) && isValidBSTSolution(root.Right, max, root.Val)
}
