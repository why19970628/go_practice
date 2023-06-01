package main

/*

这道题目要求左叶子之和，其实是比较绕的，因为不能判断本节点是不是左叶子节点。

此时就要通过节点的父节点来判断其左孩子是不是左叶子了。

平时我们解二叉树的题目时，已经习惯了通过节点的左右孩子判断本节点的属性，而本题我们要通过节点的父节点判断本节点的属性。

希望通过这道题目，可以扩展大家对二叉树的解题思路。

本题迭代法使用前中后序都是可以的，只要把左叶子节点统计出来，就可以了，
那么参考文章 二叉树：听说递归能做的，栈也能做！ 和二叉树：迭代法统一写法中的写法，可以写出一个前序遍历的迭代法。

*/

// 递归的遍历顺序为 后序遍历（左右中），是因为要通过递归函数的返回值来累加求取左叶子数值之和。
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left) // 左

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 左子树就是一个左叶子的情况
		leftValue = root.Left.Val
	}

	rightValue := sumOfLeftLeaves(root.Right) // 右

	sum := leftValue + rightValue // 中
	return sum

}

// 迭代法(前序遍历 中左右)
func sumOfLeftLeavesV22(root *TreeNode) int {
	st := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	st = append(st, root)
	result := 0

	for len(st) > 0 {
		node := st[len(st)-1] // 取最后一个值
		st = st[:len(st)-1]   // 去除最后一个值
		// 左节点是否是叶子结点
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}

	return result
}
