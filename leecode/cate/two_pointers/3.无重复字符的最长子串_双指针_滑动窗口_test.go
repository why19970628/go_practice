package main

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。


示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/

//作者：tangweiqun
//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/jian-dan-yi-dong-javac-pythonjshua-dong-bff20/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func f1_暴力法() {
	s := "bbbbb"
	r := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			res := s[i:j]
			if hasSameStr(res) {
				continue
			}
			if len(res) > r {
				r = len(res)
			}
		}
	}
	fmt.Println(r)
}

func hasSameStr(s string) bool {
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return true
		}
		m[s[i]] = struct{}{}
	}
	return false
}

func lengthOfLongestSubstringV3(s string) int {
	l, r := 0, 0
	resp := 0
	mp := make(map[byte]int)
	for r < len(s) {
		rVal := s[r]
		mp[rVal]++
		r++
		for mp[rVal] > 1 {
			if _, ok := mp[s[l]]; ok {
				mp[s[l]]--
			}
			l++
		}
		resp = max2(resp, r-l)
	}
	return resp
}

func lengthOfLongestSubstringV4(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	l, r := 0, 0
	resp := 0
	mp := make(map[byte]int)
	for r < len(s) {
		rv := s[r]
		mp[rv]++
		r++
		for mp[rv] > 1 {
			if _, ok := mp[s[l]]; ok {
				mp[s[l]]--
			}
			l++
		}
		resp = max2(resp, r-l)
	}
	fmt.Println("resp", resp)
	return resp
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestLengthOfLongestSubstringV2(t *testing.T) {
	// 3
	fmt.Println(lengthOfLongestSubstringV4("au"))
}
