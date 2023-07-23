package linked_List

/*
题意： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next == nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 环内相遇
		if slow == fast {
			// 移动head右移
			// 这就意味着，从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点，
			// 那么当这两个指针相遇的时候就是 环形入口的节点。
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

// 基于map
func detectCycleMap(head *ListNode) *ListNode {
	mp := make(map[*ListNode]*ListNode)
	for head != nil {
		if v, ok := mp[head.Next]; ok {
			return v
		}
		mp[head] = head
		head = head.Next
	}
	return nil
}
