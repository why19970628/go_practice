package main

import "fmt"

type ListNode2 struct {
	Value int
	Next  *ListNode2
}

func removeElements(head *ListNode2, val int) *ListNode2 {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Value == val {
		return head.Next
	}
	return head
}

func main() {
	a := &ListNode2{
		Value: 1,
		Next: &ListNode2{
			Value: 2,
			Next: &ListNode2{
				Value: 3,
				Next:  nil,
			},
		},
	}
	new_list := removeElements(a, 1)
	PrintListNode(new_list)
}

func PrintListNode(l *ListNode2) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}
