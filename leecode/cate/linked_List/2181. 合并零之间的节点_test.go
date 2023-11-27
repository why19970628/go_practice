package linked_List

import (
	"testing"
)

/*
给你一个链表的头节点 head ，该链表包含由 0 分隔开的一连串整数。链表的 开端 和 末尾 的节点都满足 Node.val == 0 。
对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。然后将所有 0 移除，修改后的链表不应该含有任何 0 。
 返回修改后链表的头节点 head 。
*/

func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	cur := head
	temp := make([]int, 0)
	sum := 0
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == 0 && sum > 0 {
			temp = append(temp, sum)
			sum = 0
		} else {
			sum += cur.Next.Val
		}
		if cur.Next.Next == nil {
			cur.Next = nil
		}
		cur = cur.Next
	}
	for i := 0; i < len(temp); i++ {
		head.Val = temp[i]
		if i == len(temp)-1 {
			break
		}
		head = head.Next
	}
	head.Next = nil
	return dummyHead.Next
}

func mergeNodesV2(head *ListNode) *ListNode {
	ans := head
	sum := 0
	for node := head.Next; node != nil; node = node.Next {
		if node.Val > 0 {
			sum += node.Val
		} else {
			ans.Next.Val = sum // 原地修改
			ans = ans.Next
			sum = 0
		}
	}
	ans.Next = nil
	return head.Next
}

func mergeNodesV3(head *ListNode) *ListNode {
	ans := head
	sum := 0
	cur := head.Next
	for cur != nil {
		if cur.Val > 0 {
			sum += cur.Val
		} else {
			ans.Next.Val = sum // 原地修改
			sum = 0
			ans = ans.Next
		}
		cur = cur.Next
	}
	printList(ans)
	//for node := head.Next; node != nil; node = node.Next {
	//	if node.Val > 0 {
	//		sum += node.Val
	//	} else {
	//		ans.Next.Val = sum // 原地修改
	//		ans = ans.Next
	//		sum = 0
	//	}
	//}
	ans.Next = nil
	return head.Next
}

func TestMergeNodes(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	//head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}
	head := &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0}}}}}
	printList(head)
	a := mergeNodesV3(head)
	printList(a)
}
