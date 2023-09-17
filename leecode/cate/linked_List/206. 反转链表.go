package linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }


pre->cur->temp
*/

//双指针
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

//双指针
// pre-> cur->temp
func reverseList2(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := head
	for cur != nil {
		temp := cur.Next

		cur.Next = pre

		// 往哪移动，就把哪的值付给他
		pre = cur
		cur = temp
	}
	return pre
}

// pre->cur->temp
func reverseList3(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

//  pre cur->temp

//  pre<-cur<-temp
func reverseList4(node *ListNode) *ListNode {
	pre := &ListNode{}
	cur := node
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// pre cur-> next

// pre<-cur<-next

func reverseList5(node *ListNode) *ListNode {
	pre := &ListNode{}
	cur := node
	for cur != nil {
		temp := cur.Next

		cur.Next = pre
		pre = cur
		cur = temp

		cur = cur.Next
	}
	return pre
}
