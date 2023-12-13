package linked_List

import "container/heap"

/*
https://leetcode.cn/problems/merge-k-sorted-lists/description/

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 { // 注意输入的 lists 可能是空的
		return nil
	}
	if m == 1 { // 无需合并，直接返回
		return lists[0]
	}
	left := mergeKLists(lists[:m/2])
	right := mergeKLists(lists[m/2:])
	return mergeTwoLists23(left, right)
}

func mergeTwoLists23(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 用哨兵节点简化代码逻辑
	cur := dummy         // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

// 最小堆, 堆定是最小元素
func mergeKListsv2(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h) // 堆化
	dummy := &ListNode{}
	cur := dummy

	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

//func (h hp) Len() int           { return len(h) }
//func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆
//func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
//func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
