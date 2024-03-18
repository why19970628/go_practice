package main

import (
	"fmt"
	"math"
)

/*
实现一个十进制数字报数程序，请按照数字从小到大的顺序返回一个整数数列，该数列从数字 1 开始，到最大的正整数 cnt 位数字结束。
*/

func countNumbers(cnt int) []int {
	end := int(math.Pow10(cnt) - 1)
	fmt.Println(end)
	resp := make([]int, end)
	for i := 1; i <= end; i++ {
		resp[i-1] = i
	}
	return resp
}

func main() {
	fmt.Println(countNumbers(2))
}
