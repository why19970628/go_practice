package main

import (
	"fmt"
	"strconv"
)

/*
https://leetcode.cn/problems/fizz-buzz/description/?envType=featured-list&envId=2ckc81c?envType=featured-list&envId=2ckc81c

给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：

answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
answer[i] == "Fizz" 如果 i 是 3 的倍数。
answer[i] == "Buzz" 如果 i 是 5 的倍数。
answer[i] == i （以字符串形式）如果上述条件全不满足。


示例 1：

输入：n = 3
输出：["1","2","Fizz"]
示例 2：

输入：n = 5
输出：["1","2","Fizz","4","Buzz"]
示例 3：

输入：n = 15
输出：["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

*/

func fizzBuzz(n int) []string {
	resp := make([]string, n+1)
	for i := 1; i < n+1; i++ {
		if (i >= 3 && i%3 == 0) && (i >= 5 && i%5 == 0) {
			resp[i] = "FizzBuzz"
		} else if i >= 3 && i%3 == 0 {
			resp[i] = "Fizz"
		} else if i >= 5 && i%5 == 0 {
			resp[i] = "Buzz"
		} else {
			resp[i] = strconv.Itoa(i)
		}
	}
	return resp[1:]
}

func main() {
	//fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
}
