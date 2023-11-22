package linked_List

/*

https://leetcode.cn/problems/reorder-list/solutions/1999276/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-u66q/
*/

// 876. 链表的中间结点
func middleNode143(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206. 反转链表
func reverseList143(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reorderList(head *ListNode) {
	mid := middleNode143(head)
	head2 := reverseList143(mid)
	for head2 != nil {
		//nxt := head.Next
		//nxt2 := head2.Next
		//head.Next = head2
		//head2.Next = nxt
		//head = nxt
		//head2 = nxt2

		head2Next := head2.Next
		next := head.Next

		head.Next = head2
		head2.Next = next
		head = next

		head2 = head2Next
	}
}

func reorderListV2(head *ListNode) {
	mid := middleNode143(head)
	head2 := reverseList143(mid)
	for head2.Next != nil {
		head2Next := head2.Next
		next := head.Next
		head.Next = head2
		head = next
		head2 = head2Next
	}
}
