package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		// 如果结果不为 0，说明 n 的最低位为 1
		if num&1 != 0 {
			count++
		}
		num = num >> 1
	}

	return count
}

// 使用&来判断一个数字是奇数还是偶数

func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		// 与操作：&
		// 1 & 1 = 1
		// 1 & 0 = 0
		// 0 & 1 = 0
		// 0 & 0 = 0

		// 使用&来判断一个数字是奇数还是偶数
		if num&1 != 0 {
			count++
		}
		// 最高位1向右移动
		num = num >> 1
	}

	return count
}

func main() {
	var num uint32 = 00000000000000000000010000000010
	fmt.Println(hammingWeight2(num))
}
