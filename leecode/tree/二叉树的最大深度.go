package main

import "fmt"

//type TreeNode struct {
//	value int        // 值
//	left  *TreeNode2 // 左子节点
//	right *TreeNode2 // 右子节点
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索

// 1.深度优先搜索和二叉树的前序遍历比较类似。
// 2.利用递归的方式不停下探树的深度。
// 3.递归的终止条件是如果节点为空就返回0。然后判断左右子树最大值同时加1来表示当前节点的高度。

func maxDepth(root *TreeNode) int {
	// 如果节点为空就不再递归下探深度
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	fmt.Println("1111", left)
	right := maxDepth(root.Right)
	fmt.Println("2222", right)

	var n int
	if left > right {
		n = left + 1
	} else {
		n = right + 1
	}
	return n
}

// 1.广度优先搜索利用迭代的方式将每一层的节点都放入到队列当中。
//2.队列出队清空进入下一层。
//3.利用一个变量来标记深度。每次进入下一次都给这个变量加1来记录深度。
func maxDepth2(root *TreeNode) int {
	// 根节点如果为0直接返回0
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0) // 创建一个队列
	queue = append(queue, root)   // 把根节点放入队列
	depth := 0                    // 声明深度变量
	for len(queue) > 0 {          // 队列里有值就一直循环
		size := len(queue) // 这里要把当前一层的队列遍历一遍全部出队
		for i := 0; i < size; i++ {
			v := queue[0]
			queue = queue[1:] // 出队
			if v.Left != nil {
				queue = append(queue, v.Left) // 如果有左子树，就左子树入队
			}
			if v.Right != nil {
				queue = append(queue, v.Right) // 如果有右子树，就右子树入队
			}
		}
		depth++ // 清出一层后给变量加1
	}
	return depth
}

// https://studygolang.com/articles/29470?fr=sidebar

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth3(root.Left)
	right := maxDepth3(root.Right)

	var n int
	if left > right {
		n = left + 1
	} else {
		n = right + 1
	}

	return n
}

func main() {
	node1 := &TreeNode{3, nil, nil}
	node1.Left = &TreeNode{9, nil, nil}
	node1.Right = &TreeNode{20, nil, nil}
	node1.Right.Left = &TreeNode{15, nil, nil}
	node1.Right.Right = &TreeNode{7, nil, nil}

	fmt.Println(maxDepth3(node1))

}
