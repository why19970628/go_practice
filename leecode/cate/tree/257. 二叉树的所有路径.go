package main

import "strconv"

/*
https://leetcode.cn/problems/binary-tree-paths/
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

*/

// 递归法：
func binaryTreePathsV1(root *TreeNode) []string {
	res := make([]string, 0)

	var f func(root *TreeNode, val string)
	f = func(root *TreeNode, val string) {
		// 当前是最后一个叶子结点
		if root.Right == nil && root.Left == nil {
			v := val + strconv.Itoa(root.Val)
			res = append(res, v)
		}
		// 非叶子节点
		val = val + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			f(root.Left, val)
		}
		if root.Right != nil {
			f(root.Right, val)
		}
	}

	f(root, "")

	return res
}
