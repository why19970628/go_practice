package linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/description/
*/
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
