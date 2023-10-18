package link_list

type ListNode struct {
	Value int
	Next  *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &ListNode{},
		size: 0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	head := l.head

	for i := 0; i <= index; i++ {
		head = head.Next
	}
	return head.Value
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}
	l.size++
	if index < 0 {
		index = 0
	}
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	add := &ListNode{
		Value: val,
		Next:  pred,
	}
	pred.Next = add
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--

	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
