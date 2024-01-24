package main

/*
https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/?envType=study-plan-v2&envId=top-100-liked

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。


解题思路
由于⼆叉搜索树有序的特性，所以中序遍历它，遍历到第 K 个数的时候就是结果
*/

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	count := int(0)
	resp := 0
	inorder230(root, k, &count, &resp)
	return resp
}

func inorder230(root *TreeNode, k int, count, resp *int) {
	if root == nil {
		return
	}
	inorder230(root.Left, k, count, resp)
	*count++
	if k == *count {
		*resp = root.Val
		return
	}
	inorder230(root.Right, k, count, resp)
	return
}
