package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.cn/problems/minimum-window-substring/

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	resp := ""
	for i := 0; i < len(s); i++ {
		resp1 := minWindowExpand(s, t, i)
		if len(resp1) == 0 {
			continue
		}
		resp = min76(resp, resp1)
	}
	return resp
}

func minWindowExpand(s string, t string, start int) string {
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}
	count := start
	l := start
	for start < len(s) {
		bt := s[start]
		if val, ok := mp[bt]; ok {
			if val == 1 {
				delete(mp, bt)
			} else {
				mp[bt]--
			}
		}
		if len(mp) == 0 {
			//fmt.Println("return: ", l, count, s[l:count+1])
			return s[l : count+1]
		}
		count++
		start++
	}
	return ""
}

func min76(a, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(a) < len(b) {
		return a
	}
	return b
}

func minWindowV2(s string, t string) string {
	var res string
	cnt := math.MaxInt32
	hashMap := make(map[byte]int)
	l := 0
	r := 0
	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}
	fmt.Println("1:", hashMap)
	for r < len(s) {
		hashMap[s[r]]--
		fmt.Println("2:", hashMap)
		for checkMpZero(hashMap) {
			fmt.Println("l", l, r)
			if r-l+1 < cnt {
				cnt = r - l + 1
				res = s[l : r+1]
			}
			hashMap[s[l]]++
			l++
		}
		r++
	}
	return res
}

func checkMpZero(hashMap map[byte]int) bool {
	for _, v := range hashMap {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
	// BANC
	fmt.Println("result:", minWindowV2("ADOBEC_ODE1_BANC", "ABC"))
	//fmt.Println("result:", minWindow("A", "AA"))
	//fmt.Println("result:", minWindowV2("aa", "aa"))
}
