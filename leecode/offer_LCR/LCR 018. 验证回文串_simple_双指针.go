package main

import (
	"fmt"
	"strings"
)

/*
https://leetcode.cn/problems/XltzEq/description/

给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。

本题中，将空字符串定义为有效的 回文串 。



示例 1:

输入: s = "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: s = "race a car"
输出: false
解释："raceacar" 不是回文串
*/

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isalnum(s[l]) {
			l++
		}
		for l < r && !isalnum(s[r]) {
			r--
		}
		//if l == r {
		//}

		//fmt.Println(string(s[l]), string(s[r]))
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}
