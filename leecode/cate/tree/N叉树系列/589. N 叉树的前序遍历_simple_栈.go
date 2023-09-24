package N叉树系列

/*

https://leetcode.cn/problems/n-ary-tree-preorder-traversal/

给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

*/

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	resp := make([]int, 0)

	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		resp = append(resp, top.Val)
		stack = stack[:len(stack)-1]
		stack = append(stack, top.Children...)
	}
	return resp
}
