package linked_List

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/merge-in-between-linked-lists/
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
*/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list1
	for i := 1; i < a; i++ {
		p1 = p1.Next
	}
	for i := 0; i < b; i++ {
		p2 = p2.Next
	}
	fmt.Println(p1.Val)
	fmt.Println(p2.Val)
	p1.Next = list2
	for p1.Next != nil {
		p1 = p1.Next
	}
	p1.Next = p2.Next
	return p1
}

func TestMergeInBetween(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}

	fmt.Println("Original List:")
	printList(head)
	// [0,1,2,1000000,1000001,1000002,5]

	mergeInBetween(head, 3, 4, &ListNode{Val: 100, Next: &ListNode{Val: 101}})
	printList(head)
}
