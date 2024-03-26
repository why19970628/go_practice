package main

import "fmt"

/*
https://leetcode.cn/problems/happy-number/description/?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。



示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false
*/

func isHappy(n int) bool {
	r := step(n)
	mp := map[int]bool{}
	for r > 0 {
		if mp[r] {
			return false
		}
		mp[r] = true

		r = step(r)
		if r == 1 {
			break
		}
		//fmt.Println(r)
	}
	if r == 1 {
		return true
	}
	return false
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(1111111))
}
