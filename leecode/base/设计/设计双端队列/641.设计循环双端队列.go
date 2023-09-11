package 设计双端队列

/*

设计实现双端队列。

实现 MyCircularDeque 类:

MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
*/

type Node struct {
	value int
	next  *Node
	pre   *Node
}
type MyCircularDeque struct {
	capacity int
	size     int
	head     *Node
	tail     *Node
}

func Constructor(k int) MyCircularDeque {
	head, tail := &Node{value: -1}, &Node{value: -1}
	head.next = tail
	tail.pre = head
	return MyCircularDeque{
		capacity: k,
		size:     0,
		head:     head,
		tail:     tail,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{value: value}

	next := this.head.next
	// head<-> node <-> tail
	node.next = next
	next.pre = node

	this.head.next = node
	node.pre = this.head
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{value: value}

	pre := this.tail.pre
	// head<-> node <-> tail
	node.next = this.tail
	this.tail.pre = node

	pre.next = node
	node.pre = pre
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	// head <-> node <-> next
	node := this.head.next

	next := node.next

	next.pre = this.head
	this.head.next = next
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	// pre <-> node <-> tail
	node := this.tail.pre
	pre := node.pre
	pre.next = this.tail
	this.tail.pre = pre
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	return this.head.next.value
}

func (this *MyCircularDeque) GetRear() int {
	return this.tail.pre.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size >= this.capacity
}
