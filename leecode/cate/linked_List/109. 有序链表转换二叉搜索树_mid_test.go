package linked_List

/*
https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
*/

func sortedListToBST(head *ListNode) *TreeNode {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildBST(arr, 0, len(arr)-1)
}

func buildBST(arr []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)>>1
	root := &TreeNode{
		Val:   arr[mid],
		Left:  nil,
		Right: nil,
	}
	root.Left = buildBST(arr, l, mid-1)
	root.Right = buildBST(arr, mid+1, r)
	return root
}
