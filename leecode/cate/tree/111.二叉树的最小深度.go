package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	quene := make([]*TreeNode, 0)
	quene = append(quene, root)
	res := 0
	for len(quene) > 0 {
		res++
		tempArr := make([]*TreeNode, 0)
		for i := 0; i < len(quene); i++ {
			node := quene[i]
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				tempArr = append(tempArr, node.Left)
			}
			if node.Right != nil {
				tempArr = append(tempArr, node.Right)
			}
		}
		quene = tempArr
	}
	return res
}
