package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/reverse-vowels-of-a-string/

给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

示例 1：

输入：s = "hello"
输出："holle"
示例 2：

输入：s = "leetcode"
输出："leotcede"
*/

func reverseVowels(s string) string {
	t := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isVowel(t[l]) {
			l++
		}
		for l < r && !isVowel(t[r]) {
			r--
		}
		t[l], t[r] = t[r], t[l]
		l++
		r--
	}
	return string(t)
}

func isVowel(b byte) bool {
	return strings.Contains("aeiouAEIOU", string(b))

}
func TestReverseVowels(t *testing.T) {
	fmt.Println(reverseVowels("leetcode"))
}
