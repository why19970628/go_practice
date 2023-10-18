package main

/*
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

那么本题是二叉搜索树，二叉搜索树是有序的，那得好好利用一下这个特点。
所以当我们从上向下去递归遍历，第一次遇到 cur节点是数值在[p, q]区间中，那么cur就是 p和q的最近公共祖先。

*/

// 递归
func lowestCommonAncestorBinarySearchTree(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < q.Val && root.Val < p.Val {
		root = lowestCommonAncestorBinarySearchTree(root.Right, p, q)
	}
	if root.Val > q.Val && root.Val > p.Val {
		root = lowestCommonAncestorBinarySearchTree(root.Left, p, q)
	}
	return root

}

// 迭代
func lowestCommonAncestorBinarySearchTreeV2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}
	return root

}
