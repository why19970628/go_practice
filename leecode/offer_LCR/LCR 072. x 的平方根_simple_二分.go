package main

import "fmt"

/*
给定一个非负整数 x ，计算并返回 x 的平方根，即实现 int sqrt(int x) 函数。

正数的平方根有两个，只输出其中的正数平方根。

如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。



示例 1:

输入: x = 4
输出: 2
示例 2:

输入: x = 8
输出: 2
解释: 8 的平方根是 2.82842...，由于小数部分将被舍去，所以返回 2

*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l := 1
	r := x
	mid := (l + r) / 2
	for l <= r {
		mid = (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if mid*mid > x {
		mid -= 1
	}
	return mid
}

func main() {
	//fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(5))
}
