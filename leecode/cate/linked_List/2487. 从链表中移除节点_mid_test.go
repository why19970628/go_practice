package linked_List

/*

https://leetcode.cn/problems/remove-nodes-from-linked-list/
给你一个链表的头节点 head 。
移除每个右侧有一个更大数值的节点。
返回修改后链表的头节点 head 。
*/

func removeNodes(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < cur.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
		cur = cur.Next
	}
	for i := 0; i < len(stack)-1; i++ {
		stack[i].Next = stack[i+1]
	}
	return stack[0]
}
