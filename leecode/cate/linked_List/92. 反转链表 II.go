package linked_List

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

func reverseLinkedList(head *ListNode) {
	var NewList *ListNode
	point := head

	for point != nil {
		next := point.Next
		point.Next = NewList
		NewList = point
		point = next
	}
}

// 1,2,3,4,5
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummyHead
	//第一步：从虚拟节点走left-1步，来到left节点前的一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 2 4
	//第二步：从pre再走right-left+1步到来right节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	//第三步：切断一个子链表（截取链表）
	leftNode := pre.Next
	rightNext := rightNode.Next

	//注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	//第四部：反转链表的子区间
	reverseLinkedList(leftNode)

	//第五步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = rightNext
	return dummyHead.Next
}
