package main

import "fmt"

/*
https://leetcode.cn/problems/power-of-three/description/?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

*/

func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

func main() {
	//fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(45))
}
