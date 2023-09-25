package main

/*
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

*/
var lac *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lac = nil
	lowestCommonAncestorDfs(root, p, q)
	return lac
}

func lowestCommonAncestorDfs(root, p, q *TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	}

	// 1.递归到叶子结点，结束递归

	lhasQ, lhasP := lowestCommonAncestorDfs(root.Left, p, q)
	rHasQ, rhasP := lowestCommonAncestorDfs(root.Right, p, q)

	hasQ := lhasQ || rHasQ || root.Val == q.Val
	hasP := rhasP || lhasP || root.Val == p.Val
	if hasQ && hasP && lac == nil {
		lac = root
	}
	return hasQ, hasP
}
