package build_tre

/*
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

3是根节点
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

preorder 前序 根左右
inorder 中序 左根右
*/

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	rootVal := preorder[0]
//	root := &TreeNode{
//		Val:   rootVal,
//		Left:  nil,
//		Right: nil,
//	}
//
//	rootIndex := 0
//	for i := 0; i < len(inorder); i++ {
//		if inorder[i] == rootVal {
//			rootIndex = i
//			break
//		}
//	}
//	root.Left = buildTree(preorder[:rootIndex+1], inorder[:rootIndex])
//	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
//	return root
//
//}
