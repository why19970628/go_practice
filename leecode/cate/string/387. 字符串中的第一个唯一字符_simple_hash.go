package main

import "fmt"

/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。



示例 1：

输入: s = "leetcode"
输出: 0
示例 2:

输入: s = "loveleetcode"
输出: 2
示例 3:

输入: s = "aabb"
输出: -1
*/

func firstUniqChar(s string) int {

	//  数量
	mp2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mp2[s[i]] += 1
	}
	for i := 0; i < len(s); i++ {
		if mp2[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
