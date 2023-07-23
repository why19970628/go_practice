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

// 设置一个虚拟头结点在进行移除节点操作：
// 那么因为单链表的特殊性，只能指向下一个节点，刚刚删除的是链表的中第二个，和第四个节点，那么如果删除的是头结点又该怎么办呢？
// 直接使用原来的链表来进行删除操作。
//设置一个虚拟头结点在进行删除操作。

// 虚拟头节点
func removeElements3(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
