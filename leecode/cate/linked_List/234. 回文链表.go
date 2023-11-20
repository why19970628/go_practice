package linked_List

/*
https://leetcode.cn/problems/palindrome-linked-list/description/
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*/

// 数组 快慢指针
func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	slow := 0
	fast := len(nums) - 1
	for slow < fast {
		if nums[slow] != nums[fast] {
			return false
		}
		slow++
		fast--
	}
	return true
}

// 反转链表
