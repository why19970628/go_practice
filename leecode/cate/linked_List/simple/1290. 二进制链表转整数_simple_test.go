package simple

import (
	"fmt"
	"testing"
)

/*
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
请你返回该链表所表示数字的 十进制值 。

*/

func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum<<1 + head.Val
		head = head.Next
	}
	return sum
}

func getDecimalValueV2(head *ListNode) int {
	decimalValue := 0
	power := 0

	// 1 1 0 1
	// 遍历链表，将每个节点的值转换为十进制数
	for head != nil {
		// 将当前节点的值左移相应位数并加到十进制值上
		fmt.Println("+:", head.Val, 1<<power, head.Val*(1<<power))
		decimalValue += head.Val * (1 << power)
		power++
		head = head.Next
	}

	return decimalValue
}

func TestGetDecimalValueV2(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}

	fmt.Println("Original List:")
	printList(head)
	newHead := getDecimalValueV2(head)
	fmt.Println(newHead)
}
