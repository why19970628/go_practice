package main

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

var (
	path []int
	//res  [][]int
)

// https://leetcode.cn/problems/path-sum-ii/solutions/172579/qian-xu-bian-li-mo-ban-liu-by-linbingyuan/
// 前序遍历
//var res [][]int
//
//func pathSum(root *TreeNode, sum int) [][]int {
//	res = [][]int{}
//	pathSumDfs(root, sum, []int{})
//	return res
//}

//func pathSumDfs(root *TreeNode, sum int, stack []int) {
//	// 每一次递归返回就是回溯
//	if root == nil {
//		return
//	}
//	stack = append(stack, root.Val)
//	if root.Left == nil && root.Right == nil {
//		if sum == root.Val {
//			r := make([]int, len(stack))
//			copy(r, stack)
//			res = append(res, r)
//		}
//	}
//	// @woniu 上面是在左右叶子都为空的条件下执行的，下面的dfs不会再执行了
//
//	pathSumDfs(root.Left, sum-root.Val, stack)
//	pathSumDfs(root.Right, sum-root.Val, stack)
//}

var resp [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	resp = make([][]int, 0)
	if root == nil {
		return resp
	}
	pathSum1_dfs(root, sum, []int{})
	return resp
}

func pathSum1_dfs(root *TreeNode, sum int, stack []int) {
	if root == nil {
		return
	}
	stack = append(stack, root.Val)
	// 根节点
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			tmp := make([]int, len(stack))
			copy(tmp, stack)
			resp = append(resp, tmp)
		}
	}
	pathSum1_dfs(root.Left, sum-root.Val, stack)
	pathSum1_dfs(root.Right, sum-root.Val, stack)
}
