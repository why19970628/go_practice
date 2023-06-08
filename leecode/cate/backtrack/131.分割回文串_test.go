package backtrack

import (
	"fmt"
	"testing"
)

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]
*/

var (
	res2     [][]string
	strSlice []string
)

func partition(s string) [][]string {
	partitionDfs(s, 0)
	return res2
}

func partitionDfs(s string, start int) {
	if start == len(s) {
		tmp := make([]string, len(strSlice))
		copy(tmp, strSlice)
		res2 = append(res2, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) {
			strSlice = append(strSlice, str)
			partitionDfs(s, i+1) // 寻找i+1为起始位置的子串
			strSlice = strSlice[:len(strSlice)-1]
		}
	}

}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func TestPartition(t *testing.T) {
	fmt.Println(partition("aabacccc"))
}
