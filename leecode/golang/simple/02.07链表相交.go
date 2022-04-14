package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for a := headA; a != nil; a = a.Next {
		vis[a] = true
	}
	for b := headB; b != nil; b = b.Next {
		if vis[b] == true {
			return b
		}
	}
	return nil
}
