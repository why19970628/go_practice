package main

import "fmt"

/*
示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
*/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := -1
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}

/*
开平方、平方根
*/

func mySqrtV2(x float64) float64 {
	left := float64(0)
	right := x
	for right-left > 1e-15 {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid
		} else if mid*mid < x {
			left = mid
		} else {
			return mid
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(16))
}
