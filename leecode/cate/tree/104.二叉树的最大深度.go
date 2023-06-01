package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 层序遍历
func maxDepthV2(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		tempArr := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				tempArr = append(tempArr, node.Left)
			}
			if node.Right != nil {
				tempArr = append(tempArr, node.Right)
			}
		}
		levl++
		queue = tempArr
	}
	return levl
}
