package main

// 递归
func countNodesV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodesV1(root.Right)
	}
	if root.Left != nil {
		res += countNodesV1(root.Left)
	}
	return res
}

// 迭代法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	quene := make([]*TreeNode, 0)
	quene = append(quene, root)
	var count int = 1
	for len(quene) > 0 {
		temp := make([]*TreeNode, 0)
		n := len(quene)
		for i := 0; i < n; i++ {
			node := quene[i]
			if node == nil {
				continue
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
				count++
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
				count++
			}
		}
		quene = temp
	}
	return count
}

func main() {

}
