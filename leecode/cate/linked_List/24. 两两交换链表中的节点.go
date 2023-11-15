package main

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
