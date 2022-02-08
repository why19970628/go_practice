package main

import "fmt"

type listNode struct {
	value int
	Next  *listNode
}

func getLength(head *listNode) (length int) {
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func remove(head *listNode, n int) *listNode {
	length := getLength(head)
	fmt.Println(length)
	dummy := &listNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		fmt.Println("i:", i)
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next

}

func remove2(head *listNode, n int) *listNode {
	length := getLength(head)

	a := &listNode{0, head}
	cp := a
	for i := 0; i < length-n; i++ {
		cp = cp.Next
	}

	cp.Next = cp.Next.Next

	return a.Next

}

func main() {
	l1 := &listNode{1, nil}
	l2 := &listNode{2, l1}
	l3 := &listNode{3, l2}
	l4 := &listNode{4, l3}
	remove(l4, 2)
}
