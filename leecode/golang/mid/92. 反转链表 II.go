package main

type listNode2 struct {
	value int
	Next  *listNode2
}

func reverseBetween(head *listNode2, left, right int) {
	dummyNode := &listNode2{value: -1, Next: head}
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

}

func reverseBetween2(head *listNode2, left, right int) *listNode2 {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := new(listNode2)
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

func main() {
	//reverseBetween()
}
