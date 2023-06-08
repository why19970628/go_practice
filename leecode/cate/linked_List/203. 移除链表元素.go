package linked_List

/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/

//在单链表操作
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// https://leetcode.cn/problems/remove-linked-list-elements/solutions/341875/203-yi-chu-lian-biao-yuan-su-you-ya-di-gui-c-yu-ya/
// 遍历
func removeElements2(head *ListNode, val int) *ListNode {
	// 优先处理 头结点匹配
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	// 再处理头结点之后的
	// 定义临时指针，并将其初始化为头节点
	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

func removeElements22(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
