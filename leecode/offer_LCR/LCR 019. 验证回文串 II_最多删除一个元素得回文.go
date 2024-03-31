package main

import "fmt"

/*

给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。

*/

func validPalindrome(s string) bool {
	isPalindrome := func(lo, hi int) bool {
		for lo < hi {
			if s[lo] != s[hi] {
				return false
			}

			lo++
			hi--
		}
		return true
	}
	mid := len(s) / 2
	n := len(s) - 1

	for i := 0; i < mid; i++ {
		fmt.Println(string(s[i]), string(s[n-i]))
		if s[i] != s[n-i] {
			return isPalindrome(i+1, n-i) || isPalindrome(i, n-i-1)
		}
	}
	return true

}

func main() {
	fmt.Println(validPalindrome("abca"))
}
