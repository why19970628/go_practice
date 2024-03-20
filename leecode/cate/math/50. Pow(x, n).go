package main

import "fmt"

/*
https://leetcode.cn/problems/powx-n/description/?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：
输入：x = 2.10000, n = 3
输出：9.26100

示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
*/
func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1 / pow(x, -n)
}

/*
分治 + 递归

经典的分治思想，每次只算一半即可，一半乘一半得最终

需要考虑一些奇数情况，一半乘一半后还得+1

额外注意
题意中的n次方，n可能是负数，遂需特判考虑一下，若为负数则先求正的，后续倒数一下即可
*/
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	resp := float64(1)

	half := pow(x, n/2)
	if n%2 == 1 {
		resp = x
	}

	return resp * half * half
}

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(34.00515, -3))
}
