package linked_List

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
/*
https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for q != nil {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i > 0 { // 连接同一层的两个相邻节点
				tmp[i-1].Next = node
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
