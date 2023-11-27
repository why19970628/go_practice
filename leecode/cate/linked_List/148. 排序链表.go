package linked_List

import "sort"

/*
https://leetcode.cn/problems/sort-list/description/
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	sort.Ints(arr)
	cur = head
	for i := 0; i < len(arr); i++ {
		cur.Val = arr[i]
		cur = cur.Next
	}
	return head
}
