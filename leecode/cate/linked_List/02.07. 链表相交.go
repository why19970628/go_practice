package linked_List

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

/* 让两个指针同时遍历两个链表，假设链表headA的结点数量为a，链表headB的结点数量为b，公共结点数量为c，那么： 链表A头结点到公共结点之前有a-c个结点 链表B头结点到公共结点之前有b-c个结点 指针A遍历完链表A之后，再开始遍历链表B，走到公共结点的时候，一共走了 a+(b-c) 步 指针B遍历完链表B之后，再开始遍历链表A，走到公共结点的时候，一共走了 b+(a-c) 步 如果有公共结点：那么就有 a+(b-c) = b+(a-c) 如果c > 0 ,那么此时c对应的结点就是公共结点， 如果c = 0 ,那么此时a b 同时指向 nil */

// 暴力
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var tmp *ListNode
	for headA != nil {
		tmp = headB
		for tmp != nil {
			if headA == tmp {
				return tmp
			}
			tmp = tmp.Next
		}
		headA = headA.Next
	}
	return nil
}

// map
func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	temp := make(map[*ListNode]int, 0)
	for headA != nil {
		temp[headA] = 1
		headA = headA.Next
	}
	for headB != nil {
		if temp[headB] == 1 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	traceAB, traceBA := headA, headB
	for traceAB != traceBA { // 相等即退出循环，返回其中一个节点
		if traceAB == nil { // 遍历完A，继续遍历B
			traceAB = headB
		} else {
			traceAB = traceAB.Next
		}
		if traceBA == nil { // 遍历完B，继续遍历A
			traceBA = headA
		} else {
			traceBA = traceBA.Next
		}
	}
	return traceAB
}
