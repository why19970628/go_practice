package main

import (
	"container/list"
	"fmt"
)

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

type CQueue struct {
	stack1 *list.List
	stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	fmt.Println("123")
}

type stack struct {
	arr1 []int
	arr2 []int
}

func (s *stack) push(val int) {
	s.arr1 = append(s.arr1, val)
}

func (s *stack) pop() int {
	if len(s.arr2) == 0 {
		for i := len(s.arr1) - 1; i >= 0; i++ {
			s.arr2 = append(s.arr2, s.arr1[i])
		}
	}
	if len(s.arr2) > 0 {
		val := s.arr2[len(s.arr2)-1]
		s.arr2 = s.arr2[:len(s.arr2)-1]
		return val
	}
	return -1
}
