package linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }


pre->cur->temp
*/

// 双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next

		// pre<-cur
		cur.Next = pre

		//--移动pre cur-- 移动到哪个位置，就取哪个位置的指针

		// 先向右移动pre
		pre = cur

		// 移动cur
		cur = temp

	}
	return pre
}

// pre <- cur -> next
func reverseListV2(head *ListNode) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// pre -> cur -> next
func reverseListV3(head *ListNode) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
