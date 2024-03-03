package linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	链接：https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/1992228/you-xie-cuo-liao-yi-ge-shi-pin-jiang-tou-plfs/
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	var (
		pre   *ListNode
		dummy = &ListNode{
			Next: head,
		}
	)
	p0 := dummy
	cur := head
	for ; k <= n; n -= k {

		// 翻转一组
		for i := 0; i < k; i++ {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		// pre 是头结点

		// <-pre cur

		// 见视频
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt

	}
	return dummy.Next
}
