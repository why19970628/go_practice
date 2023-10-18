package main

/*
https://leetcode.cn/problems/delete-node-in-a-bst/

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。



*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1.找到节点删除的位置
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 2.叶子节点，直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil {
			// 3.非叶子结点，找到最大值，删除最大值的节点
			root.Val = findMax(root.Left)
			root.Left = deleteNode(root.Left, root.Val)
		} else if root.Right != nil {
			root.Val = findMin(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findMax(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func findMin(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
