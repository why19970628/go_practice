package N叉树系列

/*
https://leetcode.cn/problems/n-ary-tree-postorder-traversal/

给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

*/

// 方法四：利用前序遍历反转
/*
在前序遍历中，我们会先遍历节点本身，然后从左向右依次先序遍历该每个以子节点为根的子树，而在后序遍历中，需要先从左到右依次遍历每个以子节点为根的子树，然后再访问根节点。

*/
func postorder(root *Node) []int {
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

	n := len(resp)
	for i := 0; i < n/2; i++ {
		resp[i], resp[n-1-i] = resp[n-1-i], resp[i]
	}
	return resp
}

/*
方法三：迭代优化

在后序遍历中，我们会先从左向右依次后序遍历每个子节点为根的子树，再遍历根节点本身。此时利用栈先进后出的原理，依次从右向左将子节点入栈，这样出栈的时候即可保证从左向右依次遍历每个子树。参考方法二的原理，可以提前将后续需要访问的节点压入栈中。





*/
func postorderV3(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	vis := map[*Node]bool{}
	for len(st) > 0 {
		node := st[len(st)-1]
		// 如果当前节点为叶子节点或者当前节点的子节点已经遍历过
		if len(node.Children) == 0 || vis[node] {
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			continue
		}
		// 栈先进后出，右边先入栈
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
		vis[node] = true
	}
	return
}

/*
dfs
复杂度分析
时间复杂度：O(m)O(m)O(m)，其中 mmm 为 NNN 叉树的节点。每个节点恰好被遍历一次。
空间复杂度：O(m)O(m)O(m)，递归过程中需要调用栈的开销，平均情况下为 O(log⁡m)O(\log m)O(logm)，最坏情况下树的深度为 m−1m-1m−1，需要的空间为 O(m−1)O(m-1)O(m−1)，因此空间复杂度为 O(m)O(m)O(m)。

*/
func postorderV2(root *Node) []int {
	var dfs func(node *Node)
	var resp []int
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			dfs(child)
		}
		resp = append(resp, root.Val)
	}
	return resp
}
