package linked_List

import (
	"fmt"
	"testing"
)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	pre := dummyHead

	for head != nil && head.Next != nil {
		// cur->2
		pre.Next = head.Next
		// 位置3
		next3 := head.Next.Next

		// 1->3
		head.Next.Next = head

		// 1->3
		head.Next = next3

		// 移动位置----

		// 1给pre 3给头结点
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		// 3为头结点
		head = next3
	}
	return head
}

/*
https://leetcode.cn/problems/swap-nodes-in-pairs/

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/
func SwapPairsV2(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		next1 := cur.Next
		next2 := cur.Next.Next
		next3 := next2
		cur.Next = next2
		next1.Next = next2.Next
		next2.Next = next1
		cur = next3.Next
	}

	return dummyHead.Next
}

func TestSwapPairsV2(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	fmt.Println("Original List:")
	printList(head)
	newHead := SwapPairsV2(head)
	printList(newHead)
}
