package main

import "fmt"

/*
https://leetcode.cn/problems/rotate-list/

旋转链表
给定链表的头节点，旋转链表，将链表每个节点往右移动k 个位置，原链表后k 个位置的节点则依次移动到链表头。
即，例如链表：1->2->3->4->5, k=2则返回链表 4->5->1->2->3

数据范围：链表中节点数满足
0 < n≤1000, 0<k<109
*/

func rotateLinkedListV2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n := getLength(head)
	k = k % n
	if k == 0 {
		return head
	}
	// 1->2->3->4->5, k=2则返回链表 4->5->1->2->3
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

func rotateLinkedList(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	// 获取链表长度
	length := getLength(head)
	k = k % length // 处理k大于链表长度的情况
	if k == 0 {
		return head
	}
	// 找到倒数第k+1个节点
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	//temp := fast
	// 1->2->3->4->5, k=2则返回链表 4->5->1->2->3
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead

}

// 获取链表长度
func getLength(head *ListNode) int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	// 旋转链表
	k := 2
	newHead := rotateLinkedList(head, k)

	fmt.Printf("Rotated List (k=%d):\n", k)
	printList(newHead)
}
