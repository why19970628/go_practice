package main

func reverseList(cur *listNode2) *listNode2 {
	var pre *listNode2
	for cur != nil {
		//获取下个节点temp
		temp := cur.Next
		//因要翻转，故cur.next=pre 下一个节点等于前一个节点
		cur.Next = pre
		//将当前节点赋值给前节点
		pre = cur
		//将temp节点赋值给当前节点，用于继续往下执行
		cur = temp

	}

	return pre
}

func main() {
	//reverseBetween()
}
