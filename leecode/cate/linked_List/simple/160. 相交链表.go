package simple

/*
https://leetcode.cn/problems/intersection-of-two-linked-lists/
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	mp := make(map[*ListNode]bool)
	for headA != nil {
		mp[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if mp[headB] {
			return headB
		} else {
			headB = headB.Next
		}
	}
	return nil
}
