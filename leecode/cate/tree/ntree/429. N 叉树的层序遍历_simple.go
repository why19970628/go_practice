package ntree

/*
https://leetcode.cn/problems/n-ary-tree-level-order-traversal/submissions/

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
*/
func nTreeLevelOrder(root *Node) [][]int {
	resp := make([][]int, 0)
	if root == nil {
		return resp
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {

		temp := make([]*Node, 0)
		r := []int{}
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			r = append(r, cur.Val)
			for j := 0; j < len(cur.Children); j++ {
				temp = append(temp, cur.Children[j])
			}
		}
		queue = temp
		resp = append(resp, r)
	}
	return resp
}
