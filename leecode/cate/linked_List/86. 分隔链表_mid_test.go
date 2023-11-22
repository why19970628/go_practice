package linked_List

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/partition-list/
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

*/

func partition(head *ListNode, x int) *ListNode {
	dummyHead1 := &ListNode{
		Val: -1,
	}
	dummyHead2 := &ListNode{
		Val: -1}
	// cur1 小于x的最后一个节点
	cur1, cur2 := dummyHead1, dummyHead2
	cur := head
	for cur != nil {
		if cur.Val >= x {
			next := cur.Next

			cur2.Next = cur
			cur2 = cur2.Next
			cur2.Next = nil

			cur = next
		} else {
			// 小于x
			next := cur.Next

			cur1.Next = cur
			cur1 = cur1.Next
			cur1.Next = nil

			cur = next
		}
	}
	printList(dummyHead1)
	printList(dummyHead2)

	if dummyHead1.Next == nil {
		return head
	} else {
		next := dummyHead1.Next
		cur1.Next = dummyHead2.Next
		return next
	}
}

func partitionV2(head *ListNode, x int) *ListNode {
	dummySmaller := &ListNode{} // 虚拟头节点，用于存储小于 x 的节点
	smaller := dummySmaller     // 指针用于追踪小于 x 的节点链表

	dummyLarger := &ListNode{} // 虚拟头节点，用于存储大于等于 x 的节点
	larger := dummyLarger      // 指针用于追踪大于等于 x 的节点链表

	cur := head // 当前遍历的节点

	for cur != nil {
		if cur.Val < x {
			smaller.Next = cur // 连接到小于 x 的链表
			smaller = smaller.Next
		} else {
			larger.Next = cur // 连接到大于等于 x 的链表
			larger = larger.Next
		}
		cur = cur.Next // 移动到下一个节点
	}

	// 将大于等于 x 的链表连接到小于 x 的链表之后
	smaller.Next = dummyLarger.Next
	// 将大于等于 x 的链表的结尾指向 nil（避免循环）
	larger.Next = nil

	// 返回连接后的链表
	return dummySmaller.Next
}
func TestPartition(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	newHead := partition(head, 3)
	printList(newHead)
}
