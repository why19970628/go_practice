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

func main() {
	var num uint32 = 00000000000000000000010000000010
	fmt.Println(hammingWeight(num))
}
