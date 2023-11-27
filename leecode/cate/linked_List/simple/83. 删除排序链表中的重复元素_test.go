package simple

import (
	"fmt"
	"testing"
)

/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	mp := make(map[int]struct{})
	cur := head
	for cur != nil && cur.Next != nil {
		mp[cur.Val] = struct{}{}
		next := cur.Next
		if _, ok := mp[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
			continue
		}
		cur = next
	}
	fmt.Println(mp)
	return head
}

func deleteDuplicates1_2(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fmt.Println("Original List:")
	printList(head)
	newHead := deleteDuplicates1_2(head)
	printList(newHead)
}
