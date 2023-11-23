package linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。
如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。
一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

*/

// dfs
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSubPathDFS(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubPathDFS(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if head == nil {
		return true
	}
	if head.Val != root.Val {
		return false
	}
	return isSubPathDFS(head.Next, root.Left) || isSubPathDFS(head.Next, root.Right)
}
