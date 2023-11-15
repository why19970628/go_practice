package main

import (
	"fmt"
	"testing"
)

/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/

func RemoveElement(node *ListNode, val int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: node,
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head.Next
}

func RemoveElementV2(node *ListNode, val int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: node,
	}

	cur := head
	for cur != nil {
		if cur.Next != nil {
			if cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		}
		cur = cur.Next

	}
	return head.Next
}

func RemoveElementV3(node *ListNode, val int) *ListNode {
	newNode := &ListNode{
		Val:  0,
		Next: node,
	}
	temp := newNode
	for newNode != nil && newNode.Next != nil {
		if newNode.Next.Val == val {
			newNode.Next = newNode.Next.Next
		} else {
			newNode = newNode.Next
		}
	}

	return temp.Next
}

func TestRemoveElements(t *testing.T) {
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
	b := RemoveElementV3(a, 5)
	fmt.Println(b.Val)
	fmt.Println(b.Next.Val)
}
