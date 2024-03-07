package main

import "fmt"

/*
https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envType=study-plan-v2&envId=leetcode-75

给你字符串 s 和整数 k 。

请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

英文中的 元音字母 为（a, e, i, o, u）。



示例 1：

输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母。
示例 2：

输入：s = "aeiou", k = 2
输出：2
解释：任意长度为 2 的子字符串都包含 2 个元音字母。
示例 3：

输入：s = "leetcode", k = 3
输出：2
解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
示例 4：

输入：s = "rhythms", k = 4
输出：0
解释：字符串 s 中不含任何元音字母。
示例 5：

输入：s = "tryhard", k = 4
输出：1
*/

func maxVowels(s string, k int) int {
	vowelsMp := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	sum := 0
	for i := 0; i < k; i++ {
		if vowelsMp[s[i]] {
			sum++
		}
	}
	maxVal := sum
	fmt.Println("m", maxVal)
	for i := k; i < len(s); i++ {
		left := s[i-k]
		right := s[i]
		if vowelsMp[right] {
			sum++
		}
		if vowelsMp[left] {
			sum--
		}
		if sum > maxVal {
			maxVal = sum
		}
	}
	return maxVal
}

func main() {
	fmt.Println(maxVowels("tryhard", 4))
}
