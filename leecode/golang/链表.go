package main

// 链表中倒数第k个结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(k int, node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	p := node
	q := node
	i := 0
	for ; p != nil; i++ {
		p = p.Next
		if i >= k-1 {
			q = q.Next
		}
	}
	if i >= k-1 {
		return q
	} else {
		return nil
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {

	//快慢指针
	fast := head
	slow := head

	//快指针先走k步
	for k > 0 {
		fast = fast.Next
		k--
	}

	//快慢指针走剩下的，直到快指针走到尾
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	//慢指针就是所求的值
	return slow.Val
}
