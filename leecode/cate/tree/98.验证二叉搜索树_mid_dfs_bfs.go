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
	return isValidBSTDFSV2(root, math.MaxInt64, math.MinInt64)
}
func isValidBSTDFSV2(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTDFSV2(root.Left, root.Val, min) && isValidBSTDFSV2(root.Right, max, root.Val)
}

func isValidBSTDFS(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTDFS(root.Left, root.Val, min) && isValidBSTDFS(root.Right, max, root.Val)
}

func isValidBSTBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	boundsQueue := [][]int{{-math.MaxInt64, math.MaxInt64}}
	for len(stack) > 0 {
		n := len(stack)
		for i := 0; i < n; i++ {
			node := stack[0]
			stack = stack[1:]
			// 记录每一个层级最小最大值
			// 在这个解决方案中，boundsQueue 中的每个元素都是一个包含两个数字的数组。这两个数字分别表示当前处理节点的允许值范围的上下界。
			// 比如，对于一个节点 node，bounds 数组 [lower, upper] 表示该节点的值必须大于 lower 并且小于 upper，否则就不符合二叉搜索树的条件。
			bounds := boundsQueue[0]
			boundsQueue = boundsQueue[1:]

			if node.Val <= bounds[0] || node.Val >= bounds[1] {
				return false
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
				boundsQueue = append(boundsQueue, []int{bounds[0], node.Val})
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
				boundsQueue = append(boundsQueue, []int{node.Val, bounds[1]})
			}
		}
	}
	return true
}

func isValidBSTV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*TreeNode
	var inorder float64 = -math.MaxFloat64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 左
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if float64(root.Val) <= inorder {
			return false
		}
		inorder = float64(root.Val)
		root = root.Right // 右
	}
	return true
}

func isValidV3(root *TreeNode) bool {
	return isValidDFSV3(root, math.MaxInt64, math.MinInt64)
}

func isValidDFSV3(root *TreeNode, maxValue, minValue int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxValue || root.Val <= minValue {
		return false
	}

	return isValidDFSV3(root.Left, root.Val, minValue) && isValidDFSV3(root.Right, maxValue, root.Val)
}
