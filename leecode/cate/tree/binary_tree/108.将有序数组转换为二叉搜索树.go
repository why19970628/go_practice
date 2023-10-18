package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { //终止条件，最后数组为空则可以返回
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{Val: nums[idx]}

	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])

	return root
}
