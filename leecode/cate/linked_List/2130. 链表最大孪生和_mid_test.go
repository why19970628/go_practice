package linked_List

/*
https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/description/


1.数组
2.快慢指针找中间节点+反转慢链表+ 遍历半个数组
*/

func pairSum(head *ListNode) int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	sum := 0
	l, r := 0, len(arr)-1
	for l < r {
		sum = max(sum, arr[l]+arr[r])
		l++
		r--
	}
	return sum
}
