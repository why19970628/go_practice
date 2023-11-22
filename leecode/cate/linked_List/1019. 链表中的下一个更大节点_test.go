package linked_List

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/next-greater-node-in-linked-list/

1019. 链表中的下一个更大节点 中等
给定一个长度为 n 的链表 head

对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。

返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设置 answer[i] = 0 。


*/

// 单调栈
func nextLargerNodes(head *ListNode) []int {
	stack := make([]int, 0)
	res := make([]int, 0)
	curList := make([]int, 0)
	count := 0
	for head != nil {
		res = append(res, 0)
		curList = append(curList, head.Val)
		for len(curList) > 0 && len(stack) > 0 && head.Val > curList[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, count)
		head = head.Next
		count++
	}
	return res
}

func TestNextLargerNodes(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	// 旋转链表
	k := 2
	newHead := nextLargerNodes(head)

	fmt.Printf("Rotated List (k=%d):\n", k)
	fmt.Println(newHead)
}
