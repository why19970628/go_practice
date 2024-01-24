package linked_List

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode.cn/problems/add-two-numbers/description/

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	a := sum(l1)
	b := sum(l2)
	c := a + b
	if c == 0 {
		return &ListNode{
			Val: 0,
		}
	}
	head := &ListNode{}
	d := &ListNode{
		Val:  0,
		Next: head,
	}
	for c > 0 {
		node := &ListNode{
			Val: c % 10,
		}
		head.Next = node
		head = node
		c = c / 10
	}
	return d.Next.Next

}

func sum(l1 *ListNode) int {
	sum1 := 0
	a := 1
	for l1 != nil {
		sum1 += l1.Val * a
		a = a * 10
		l1 = l1.Next
	}
	return sum1
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	a := sumv2(l1)
	b := sumv2(l2)
	fmt.Println("a", a)
	fmt.Println("b", b)

	c := a + b
	fmt.Println("c", c)
	cs := strconv.Itoa(c)
	head := &ListNode{}
	d := &ListNode{
		Val:  0,
		Next: head,
	}
	for i := len(cs) - 1; i >= 0; i-- {
		d, _ := strconv.Atoi(string(cs[i]))
		fmt.Println("d", d)
		code := &ListNode{
			Val:  d,
			Next: nil,
		}
		head.Next = code
		head = code
	}
	return d.Next.Next
}

func sumv2(l1 *ListNode) int {
	s := []string{}
	for l1 != nil {
		s = append(s, strconv.Itoa(l1.Val))
		l1 = l1.Next
	}
	a := ""
	for i := len(s) - 1; i >= 0; i-- {
		a = fmt.Sprintf("%s%s", a, s[i])
	}
	res, _ := strconv.Atoi(a)
	return res

}

func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	t := 0
	cur := dummyHead
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{
			Val:  t % 10,
			Next: nil,
		}
		cur = cur.Next
		t = t / 10
	}
	return dummyHead.Next
}

func addTwoNumbersV4(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	t := 0
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{
			Val: t % 10,
		}
		cur = cur.Next
		t = t / 10
	}
	return dummyHead.Next
}

func TestAddTwoNumbers(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	head2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	//head := &ListNode{Val: 0}
	//head2 := &ListNode{Val: 0}

	fmt.Println("Original List:")
	printList(head)
	printList(head2)

	newHead := addTwoNumbersV4(head, head2)
	printList(newHead)
}
