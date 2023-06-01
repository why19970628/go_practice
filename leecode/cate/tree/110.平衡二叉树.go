package main

/*
https://leetcode.cn/problems/balanced-binary-tree/solutions/866942/dai-ma-sui-xiang-lu-dai-ni-xue-tou-er-ch-x3bv/
通过本题可以了解求二叉树深度 和 二叉树高度的差异，求深度适合用前序遍历（中左右），而求高度适合用后序遍历（左右中）。

在104.二叉树的最大深度中, 我们可以使用层序遍历来求深度，但是就不能直接用层序遍历来求高度了，这就体现出求高度和求深度的不同。
*/

func isBalanced(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getHeight(root.Left), getHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}

/*

作者：灵茶山艾府
链接：https://leetcode.cn/problems/balanced-binary-tree/solutions/2015068/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-c3wj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 要求比较高度，必然是要后序遍历 (左右中)。
func getHeightV2(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := getHeightV2(node.Left)
	if leftH == -1 {
		return -1 // 提前退出，不再递归
	}
	rightH := getHeightV2(node.Right)
	if rightH == -1 || abs(leftH-rightH) > 1 {
		return -1
	}
	return max(leftH, rightH) + 1
}

func isBalanced2(root *TreeNode) bool {
	return getHeightV2(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
