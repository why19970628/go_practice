package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	fmt.Println(dp, m, n)
	return dp[m][n]
}

func longestCommonSubsequenceV2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Println(dp, m, n)
	return dp[m][n]
}

/*
给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。
*/
func longestCommonSubstring(str1, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	maxLen := 0
	endIndex := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					endIndex = i
				}
			}
		}
	}
	return str1[endIndex-maxLen+1 : endIndex+1]
}

func main() {
	str1 := "abcdefg"
	str2 := "bcde"
	fmt.Println(longestCommonSubstring(str1, str2)) // 输出 "bcde"

	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
