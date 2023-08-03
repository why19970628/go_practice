package main

import (
	"testing"
)

/*
本题要找出树的最后一行的最左边的值。此时大家应该想起用层序遍历是非常简单的了，反而用递归的话会比较难一点。

*/

var depth int  // 全局变量 最大深度
var resInt int // 记录最终结果
func findBottomLeftValueV1(root *TreeNode) int {
	depth, resInt = 0, 0 // 初始化
	findBottomLeftValueDFS(root, 1)
	return resInt
}

func findBottomLeftValueDFS(root *TreeNode, d int) {
	if root == nil {
		return
	}
	// 需要先遍历左边，左边如果有值，同层的右边不会更新结果
	// 如果是根节点，并且当前深度比最大深度大，则取值
	if root.Left == nil && root.Right == nil && depth < d {
		depth = d
		resInt = root.Val
	}
	findBottomLeftValueDFS(root.Left, d+1) // 隐藏回溯
	findBottomLeftValueDFS(root.Right, d+1)
}

// 层序遍历 BFS
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	quene := make([]*TreeNode, 0)
	quene = append(quene, root)
	var res int
	for len(quene) > 0 {
		tempArr := make([]*TreeNode, 0)
		for i := 0; i < len(quene); i++ {
			node := quene[i]
			if node.Left != nil {
				tempArr = append(tempArr, node.Left)
			}
			if node.Right != nil {
				tempArr = append(tempArr, node.Right)
			}
			// 每一层取最左节点的值
			if i == 0 {
				res = node.Val
			}
		}
		quene = tempArr
	}
	return res
}

func Test2(t *testing.T) {
	node1 := &TreeNode{2, nil, nil}
	node1.Left = &TreeNode{1, nil, nil}
	node1.Right = &TreeNode{3, nil, nil}
	t.Log(findBottomLeftValue(node1))
}
