package main

import (
	"fmt"
	"strconv"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。



示例 1：

输入：num1 = "11", num2 = "123"
输出："134"
示例 2：

输入：num1 = "456", num2 = "77"
输出："533"
示例 3：

输入：num1 = "0", num2 = "0"
输出："0"
*/

func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)

	n := max(n1, n2)
	resp := []int{}
	carry := 0
	for i := 0; i < n; i++ {

		var a, b int
		if i < n1 {
			a = int(num1[n1-i-1] - '0')
		}
		if i < n2 {
			b = int(num2[n2-i-1] - '0')
		}
		resp = append(resp, (a+b+carry)%10)
		carry = (a + b + carry) / 10
	}
	if carry > 0 {
		resp = append(resp, 1)
	}
	r := ""
	for i := len(resp) - 1; i >= 0; i-- {
		r += strconv.Itoa(resp[i])
	}
	return r
}

func main() {
	fmt.Println(addStrings("456", "77")) // 533
	fmt.Println(addStrings("9", "1"))    // 533
}
