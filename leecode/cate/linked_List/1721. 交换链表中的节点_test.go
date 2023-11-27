package linked_List

import (
	"testing"
)

/*
给你链表的头节点 head 和一个整数 k 。
交换 链表正数第 k 个节点和倒数第 k 个节点的值后，返回链表的头节点（链表 从 1 开始索引）。


*/

func swapNodes(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	fast, slow := head, head
	temp1 := &ListNode{}
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	temp1 = fast
	a := fast.Val
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	b := slow.Val
	temp1.Val = b
	slow.Val = a
	printList(dummyHead)
	return dummyHead.Next
}

func TestSwapNodes(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	//head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	printList(head)
	swapNodes(head, 2)
	printList(head)
}
