package main

/*
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
https://leetcode.cn/problems/symmetric-tree/
*/

// 递归
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)

}

//func dfs(left *TreeNode, right *TreeNode) bool {
//	if left == nil && right == nil {
//		return true
//	}
//	if left == nil || right == nil {
//		return false
//	}
//	if left.Val != right.Val {
//		return false
//	}
//	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
//}

// 迭代
func isSymmetricV2(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
