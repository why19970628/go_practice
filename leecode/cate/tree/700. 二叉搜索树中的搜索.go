package main

/*
https://leetcode.cn/problems/search-in-a-binary-search-tree/description/

给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。

*/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
