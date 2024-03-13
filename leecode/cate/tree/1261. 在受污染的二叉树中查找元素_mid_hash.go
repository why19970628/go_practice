package main

/*
https://leetcode.cn/problems/find-elements-in-a-contaminated-binary-tree/description/

给出一个满足下述规则的二叉树：

root.val == 0
如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。

请你先还原二叉树，然后实现 FindElements 类：

FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
*/

type FindElements struct {
	mp map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{
		mp: make(map[int]bool),
	}
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		f.mp[target] = true
		dfs(root.Left, target*2+1)
		dfs(root.Right, target*2+2)
	}
	dfs(root, 0)
	return f
}

func (this *FindElements) Find(target int) bool {
	return this.mp[target]
}
