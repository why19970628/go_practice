package main

/*

112.路径总和I    https://leetcode.cn/problems/path-sum/description/
113.路径总和II   https://leetcode.cn/problems/path-sum-ii/description/
437.路径总和III  https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=leetcode-75


给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

func pathSum3(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	resp := rootSum(root, targetSum)
	resp += pathSum3(root.Left, targetSum)
	resp += pathSum3(root.Right, targetSum)
	return resp
}

func rootSum(root *TreeNode, targetSum int) (resp int) {
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		resp++
	}
	resp += rootSum(root.Left, targetSum-root.Val)
	resp += rootSum(root.Right, targetSum-root.Val)
	return
}

func main() {

}
