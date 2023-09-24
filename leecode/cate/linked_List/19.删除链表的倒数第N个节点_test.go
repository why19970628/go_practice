package linked_List

import (
	"fmt"
	"testing"
)

/*
双指针的经典应用，如果要删除倒数第n个节点，让fast移动n步，然后让fast和slow同时移动，直到fast指向链表末尾。删掉slow所指向的节点就可以了。

*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	fast := head
	slow := dummyHead
	i := 1

	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}
	// 删除第N个节点
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func RemoveNthFromEndV1(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	fast := head
	slow := newHead
	i := 0
	for fast != nil {
		fast = fast.Next
		i++
		if i > n {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	temp := newHead

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		for i := 0; i < n; i++ {
			fast = fast.Next
		}
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}

func TestRemoveNthFromEndV1(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	b := removeNthFromEnd2(a, 2)
	for b != nil {
		fmt.Println("node", b.Val)
		b = b.Next
	}
}

func RemoveNthFromEndV2(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	fast := head
	slow := newHead
	count := 0
	for fast != nil {
		fast = fast.Next
		count++
	}
	fmt.Println(count)
	curCount := 0
	for slow != nil && slow.Next != nil {
		fmt.Println(count - curCount)
		if count-curCount == n {
			slow.Next = slow.Next.Next
		}
		curCount++
		slow = slow.Next
	}
	return newHead.Next
}

func TestRemoveNthFromEndV2(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	b := RemoveNthFromEndV2(a, 2)
	for b != nil {
		fmt.Println("node", b.Val)
		b = b.Next
	}
}

func RemoveNthFromStart(head *ListNode, n int) *ListNode {
	newhead := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := newhead

	count := 0
	for cur != nil && cur.Next != nil {
		count++
		fmt.Println(count, count == n)
		if count == n {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return newhead.Next
}

func TestRemoveNthFromStart(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	b := RemoveNthFromStart(a, 1)
	fmt.Println(b.Val)
	fmt.Println(b.Next.Val)
}
