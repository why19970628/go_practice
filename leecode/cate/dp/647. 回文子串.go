package main

/*
https://leetcode.cn/problems/palindromic-substrings/description/
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例 1：

输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
提示：输入的字符串长度不会超过 1000 。

#
*/

func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				result++
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func countSubstringsDP(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}

// 双指针
func countSubstringsV3(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		ans += extend(s, i, i, len(s))
		ans += extend(s, i, i+1, len(s))
	}
	return ans
}

func extend(s string, i, j int, n int) int {
	ans := 0
	for i >= 0 && j < n && s[i] == s[j] {
		ans++
		i--
		j++
	}
	return ans
}

// https://leetcode.cn/problems/palindromic-substrings/solutions/380130/shou-hua-tu-jie-dong-tai-gui-hua-si-lu-by-hyj8/
