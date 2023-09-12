package main

/*

 */

// 后序展开为右链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left, right := root.Left, root.Right
	root.Left = nil
	root.Right = left
	// 找到z之前左节点的根节点，连接右节点
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
