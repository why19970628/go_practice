package main

import "fmt"

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
*/
func findAnagrams(s string, p string) []int {
	n := len(p)
	cnt := [26]int{}
	for _, i := range p {
		cnt[i-'a']++
	}
	resp := make([]int, 0)
	for i := 0; i < len(s)-n+1; i++ {
		cur := s[i : i+n]
		curCnt := [26]int{}
		for _, i2 := range cur {
			curCnt[i2-'a']++
		}
		if cnt == curCnt {
			resp = append(resp, i)
		}
	}
	return resp
}

func main() {
	fmt.Println("sortStr", findAnagrams("abab", "ab"))
}
