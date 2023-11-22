package linked_List

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
https://leetcode.cn/problems/linked-list-random-node/description/
382. 链表随机节点

给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。

实现 Solution 类：

Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。

*/

type Solution struct {
	mp map[*ListNode]*ListNode
}

func Constructor(head *ListNode) Solution {
	mp := map[*ListNode]*ListNode{
		head: head,
	}
	for head != nil {
		mp[head] = head
		head = head.Next
	}
	return Solution{
		mp: mp,
	}
}

func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(this.mp))
	index := 0
	for node, _ := range this.mp {
		if index == n {
			return node.Val
		}
		index++
	}
	return -1
}
func TestGetRandomListNode(t *testing.T) {
	// 创建一个链表: 1->2->3->4->5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Println("Original List:")
	printList(head)

	c := Constructor(head)
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())

}
