package linked_List

import (
	"fmt"
	"testing"
)

/*
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。

每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。

这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。

返回一个由上述 k 部分组成的数组。



输入：head = [1,2,3], k = 5
输出：[[1],[2],[3],[],[]]
解释：
第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。

输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
输出：[[1,2,3,4],[5,6,7],[8,9,10]]
解释：
输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。前面部分的长度大于等于后面部分的长度。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	length := getLength2(head)
	res := make([]*ListNode, k)
	if length == 0 {
		return res
	}

	offset := length / k
	remainder := length % k

	fmt.Println("length", length, offset)
	for i := 0; i < k; i++ {
		res[i] = head
		for j := 0; j < offset-1; j++ {
			head = head.Next
		}
		if remainder > 0 && head != nil && offset > 0 {
			head = head.Next
			remainder--
		}
		if head != nil {
			head.Next, head = nil, head.Next
		}
	}

	return res
}

// 获取链表长度
func getLength2(head *ListNode) int {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	return length
}

func TestSplitListToParts(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	// 旋转链表
	k := 2
	newHead := splitListToParts(head, 3)

	fmt.Printf("Rotated List (k=%d):\n", k)
	for _, node := range newHead {
		printList(node)
	}
}
