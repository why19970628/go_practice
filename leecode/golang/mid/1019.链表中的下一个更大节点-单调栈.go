package main

// 给定一个长度为n的链表head
//
//对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
//
//返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设置answer[i] = 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/next-greater-node-in-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	// 遍历链表切片
	res := list2Slice(head)
	if res == nil {
		return nil
	}
	ans := make([]int, 0, len(res))
	// 单调栈
	tmp := make([]int, 0)
	for index, value := range res {
		for len(ans) > 0 && value > res[tmp[len(tmp)-1]] {
			rightIndex := tmp[len(tmp)-1]
			ans[rightIndex] = value
			// 最右元素出栈
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, index)
	}

	return tmp
}

func list2Slice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
