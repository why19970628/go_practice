package linked_List

type SingleNode struct {
	Next *SingleNode
	Val  int
}

type SingleLinkedList struct {
	head *SingleNode
	size int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{head: new(SingleNode), size: 0}
}

func (l *SingleLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	// 让cur等于真正头节点
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *SingleLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *SingleLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *SingleLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	l.size++
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &SingleNode{pred.Next, val}
	pred.Next = toAdd
}

func (l *SingleLinkedList) DeleteAtIndex(index int) {
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

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
