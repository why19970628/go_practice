package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

438 一样
*/

func groupAnagrams(strs []string) [][]string {
	resp := make([][]string, 0)
	mp := make(map[string][]string)
	for _, str := range strs {
		fmt.Println(str)
		s := sortStr(str)
		mp[s] = append(mp[s], str)
	}
	for _, strs := range mp {
		resp = append(resp, strs)
	}
	return resp

}

func sortStr(s string) string {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func groupAnagramsV2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, i := range str {
			cnt[i-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	resp := make([][]string, 0, len(mp))
	for _, strings := range mp {
		resp = append(resp, strings)
	}
	return resp

}

func main() {
	fmt.Println("sortStr", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
