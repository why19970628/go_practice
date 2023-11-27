package simple

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil && cur.Next.Next != nil && cur.Next.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			cur.Next = cur.Next.Next.Next
		}
	}
	return head
}

func deleteDuplicates21(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if val == cur.Next.Next.Val {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

func TestDeleteDuplicates2(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	fmt.Println("Original List:")
	printList(head)
	newHead := deleteDuplicates21(head)
	printList(newHead)
}
