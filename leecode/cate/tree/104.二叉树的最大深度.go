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
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levl++
	}
	return levl
}

func maxDepthDFS(root *TreeNode) (resp int) {
	if root == nil {
		return
	}
	var dfs func(node *TreeNode, cur int) int
	cur := 0
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur++
		resp = max(resp, cur)
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		return resp
	}
	dfs(root, cur)
	return
}
