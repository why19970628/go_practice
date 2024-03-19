package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.cn/problems/reverse-integer/description/?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

*/

func reverse(x int) int {
	resp := 0

	for x != 0 {
		// 当我们试图将一个整数 reversed 乘以 10 并加上一个数字时，我们需要确保结果不会超过 math.MaxInt32 或小于 math.MinInt32。由于我们正在进行整数除法，因此在每次迭代时 reversed 只会增加一个数字的倍数。
		// 提前判断
		if resp < math.MinInt32/10 || resp > math.MaxInt32/10 {
			return 0
		}

		digit := x % 10

		resp = resp*10 + digit

		x = x / 10
	}
	return resp
}

func main() {
	fmt.Println(reverse(123))
}
