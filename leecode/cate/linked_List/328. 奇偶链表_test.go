package linked_List

import (
	"fmt"
	"testing"
)

/*
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

输入: head = [1,2,3,4,5]
输出: [1,3,5,2,4]
*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummnyhead1 := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummnyhead2 := &ListNode{
		Val:  -1,
		Next: nil,
	}

	cur1, cur2 := dummnyhead1, dummnyhead2
	cur := head
	for i := 0; cur != nil; i++ {
		// 偶数
		if i%2 == 0 {
			next := cur.Next

			cur1.Next = cur
			cur1 = cur1.Next
			cur1.Next = nil

			cur = next
		} else {
			next := cur.Next
			cur2.Next = cur
			cur2 = cur2.Next
			cur2.Next = nil
			cur = next
		}
	}
	printList(cur1)
	printList(cur2)
	fmt.Println("---")
	if dummnyhead2.Next == nil {
		return dummnyhead1.Next
	} else {
		cur1.Next = dummnyhead2.Next
		return dummnyhead1.Next
	}
}

func oddEvenListV2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead1 := &ListNode{
		Val: -1,
	}
	dummyHead2 := &ListNode{
		Val: -1,
	}
	cur1, cur2 := dummyHead1, dummyHead2
	cur := head
	index := 0
	for cur != nil {
		// 偶数
		if index%2 == 1 {
			next := cur.Next
			cur2.Next = cur
			cur2 = cur2.Next
			cur2.Next = nil
			cur = next
		} else {
			next := cur.Next
			cur1.Next = cur
			cur1 = cur1.Next
			cur1.Next = nil
			cur = next
		}
		index++
	}
	printList(dummyHead1)
	printList(dummyHead2)
	// 偶数
	if dummyHead2.Next == nil {
		return dummyHead1.Next
	} else {
		cur1.Next = dummyHead2.Next
		return dummyHead1.Next
	}
}

func TestOddEvenList(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	newHead := oddEvenListV2(head)
	printList(newHead)
}
