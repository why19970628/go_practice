package main

import "fmt"

/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2

示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1

说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
func strStr(a, b string) bool {
	abt := []byte(a)
	bbt := []byte(b)

	an := len(abt)
	bn := len(bbt)

	for i := 0; i < an; i++ {
		fmt.Println("i", i)
		var is bool
		for j := 0; j < bn; j++ {
			fmt.Println("---", j, string(abt[i+j]), string(bbt[j]))
			if abt[i+j] != bbt[j] {
				break
			}
			if j == bn-1 {
				is = true
			}
		}
		if is {
			return true
		}

	}
	return false
}

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
*/

func strStrV2(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n1 < n2 {
		return -1
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			// 剩余的不够了，直接返回
			if i+j > n1-1 {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
			if j == n2-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("leetcode", "leet"))
	fmt.Println(strStrV2("leetcode", "leet"))
	fmt.Println(strStrV2("sadbutsad", "sad"))
}
