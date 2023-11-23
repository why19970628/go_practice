package linked_List

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/aMhZSa/

给定一个链表的 头节点 head ，请判断其是否为回文链表。
如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。

找到链表的中间节点，反转后半段链表
*/
func isPalindromeListNode(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	count := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		count++
	}
	//printList(head)
	printList(slow)
	slow = reverseList143(slow)

	for i := 0; i < count; i++ {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func isPalindromeListNodeV2(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	l, r := 0, len(list)-1
	for l < r {
		if list[l] != list[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestIsPalindromeListNode(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}

	fmt.Println("Original List:")
	printList(head)
	newHead := isPalindromeListNode(head)
	fmt.Println(newHead)
}
