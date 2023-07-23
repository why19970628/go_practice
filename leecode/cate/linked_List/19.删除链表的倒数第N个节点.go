package linked_List

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
