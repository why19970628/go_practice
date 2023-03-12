package main

import "fmt"

type TreeNode111 struct {
	Val   int
	Left  *TreeNode111
	Right *TreeNode111
}

func maxDepth(root *TreeNode111) int {
	i := 0
	j := 0
	if root == nil {
		return 0
	}
	i = maxDepth(root.Left)
	j = maxDepth(root.Right)
	if i < j {
		i, j = j, i
	}
	return i + 1
}

//func maxDepth2(root *TreeNode) int {
//	var ans int = 0
//	var f func(root *TreeNode, cnt int)
//	f = func(root *TreeNode, cnt int) {
//		if root == nil {
//			return
//		}
//		cnt += 1
//		if cnt > ans {
//			ans = cnt
//		}
//		f(root.Left, cnt)
//		f(root.Right, cnt)
//	}
//	f(root, 0)
//	return ans
//}

func main() {
	node1 := &TreeNode111{3, nil, nil}
	node1.Left = &TreeNode111{9, nil, nil}
	node1.Right = &TreeNode111{20, nil, nil}
	node1.Right.Left = &TreeNode111{15, nil, nil}
	node1.Right.Right = &TreeNode111{7, nil, nil}
	fmt.Println(maxDepth(node1))
}
