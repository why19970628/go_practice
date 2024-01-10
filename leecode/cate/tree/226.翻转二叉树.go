package main

// / 递归版本的前序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 递归版本的中序遍历
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree2(root.Left)
	root.Left, root.Right = root.Right, root.Left
	invertTree2(root.Right)
	return root
}

// 递归版本的后序遍历
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 层序遍历 把放入切片前，把左右两边指针交换位置即可
func invertTree_BFS(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		var newArr []*TreeNode
		for _, v := range curLevel {
			v.Left, v.Right = v.Right, v.Left
			if v.Right != nil {
				newArr = append(newArr, v.Right)
			}
			if v.Left != nil {
				newArr = append(newArr, v.Left)
			}
		}
		curLevel = newArr
	}
	return root
}

func invertTree_BFS2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		v := curLevel[0]
		curLevel = curLevel[1:]
		if v.Left != nil {
			curLevel = append(curLevel, v.Left)
		}
		if v.Right != nil {
			curLevel = append(curLevel, v.Right)
		}
		v.Left, v.Right = v.Right, v.Left
	}
	return root
}
