package main

/*
https://leetcode.cn/problems/find-words-that-can-be-formed-by-characters/
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。
注意：每次拼写（指拼写词汇表中的一个单词）时，chars 中的每个字母都只能用一次。
返回词汇表 words 中你掌握的所有单词的 长度之和。

示例 1：
输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6
解释：
可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。

示例 2：
输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
输出：10
解释：
可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。
*/
func countCharacters(words []string, chars string) int {
	arr := make([]int, 'z'+1)
	for i := 0; i < len(chars); i++ {
		arr[chars[i]]++
	}
	resp := 0
	for _, word := range words {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		var flag bool
		for i := 0; i < len(word); i++ {
			if newArr[word[i]] <= 0 {
				break
			} else {
				newArr[word[i]]--
			}
			if i == len(word)-1 {
				flag = true
			}
		}
		if flag {
			resp += len(word)
		}
	}

	return resp
}
