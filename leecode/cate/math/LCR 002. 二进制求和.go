package main

import (
	"fmt"
	"strconv"
)

/*
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
输入为 非空 字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "10"
输出: "101"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/

func addBinary(a string, b string) string {
	n1, n2 := len(a), len(b)
	n := max(n1, n2)

	r := []int{}

	carry := 0
	for i := 0; i < n; i++ {
		x, y := 0, 0
		if i < n1 {
			x = int(a[n1-i-1] - '0')
		}
		if i < n2 {
			y = int(b[n2-i-1] - '0')
		}
		fmt.Println(x, y)
		cur := (carry + x + y) % 2

		carry = (carry + x + y) / 2
		r = append(r, cur)
	}
	if carry > 0 {
		r = append(r, 1)
	}
	resp := ""
	for i := len(r) - 1; i >= 0; i-- {
		resp += strconv.Itoa(r[i])
	}

	return resp
}

func main() {
	fmt.Println(addBinary("11", "10"))
}
