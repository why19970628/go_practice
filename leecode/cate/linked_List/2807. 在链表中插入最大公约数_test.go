package linked_List

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/
*/
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		nxt := cur.Next

		cVal := cur.Val
		nVal := cur.Next.Val
		r := divisor(cVal, nVal)
		fmt.Println("r:", r)
		cur.Next = &ListNode{Val: r, Next: nxt}
		cur = nxt
	}
	return head
}

func divisorV2(min, max int) int {
	resp := 0
	a := max % min
	if a != 0 {
		resp = divisorV2(a, min)
	} else {
		resp = min
	}
	return resp

}

func TestDivisorV2(t *testing.T) {
	fmt.Println(divisorV2(5, 12))
}

// / 获取最大公约数
func divisor(min, max int) (maxDivisor int) {
	if min > max {
		min, max = max, min
	}
	// 用小数除大数取余
	complement := max % min
	// 余数不为零，将余数作为小数，除数作为大数，递归调用自己
	if complement != 0 {
		maxDivisor = divisor(complement, min)
	} else {
		//余数为零，除数则为最大公约数
		maxDivisor = min
	}
	return
}
