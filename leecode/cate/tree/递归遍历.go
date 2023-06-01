package main

// 前序遍历 根 ---> 左 --->右
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraverse(node *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(node)
	return res
}

// 中序遍历 左 ---> 根 --->右
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 因此，树从小到大的遍历顺序是中序遍历。中序遍历会按照节点值的升序访问树中的节点。
func inorderTraverse(node *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(node)
	return res
}

// 后序遍历 左 ----> 右 ---> 根
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postTraverse(node *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(node)
	return res
}
