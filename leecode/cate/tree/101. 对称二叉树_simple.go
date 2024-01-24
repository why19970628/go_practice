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

func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricDfs(root.Left, root.Right)
}

func isSymmetricDfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricDfs(left.Left, right.Right) && isSymmetricDfs(left.Right, right.Left)
}

// 层序遍历
func isSymmetricBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var quene []*TreeNode
	quene = append(quene, root.Left, root.Right)
	for len(quene) > 0 {
		n := len(quene)
		for i := 0; i < n/2; i++ {
			left := quene[0]
			right := quene[1]
			quene = quene[2:]
			if left == nil && right == nil {
				continue
			}
			if left == nil || right == nil {
				return false
			}
			if left.Val != right.Val {
				return false
			}
			quene = append(quene, left.Left, right.Right, left.Right, right.Left)
		}
	}
	return true
}

func isSymmetricV3(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(left, right *TreeNode) bool

	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left != nil && right == nil {
			return false
		}
		if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}
