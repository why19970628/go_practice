package main

import "fmt"

func main() {
	//f1_暴力法()
	fmt.Println(lengthOfLongestSubstring2("dvdf"))
}

/*
   r
abcadb
l

a 1
b 1
c 1
*/
func lengthOfLongestSubstring(s string) int {
	resp := 0
	left, right := 0, 0
	mp := make(map[byte]int)

	for right < len(s) {
		r := s[right]
		mp[r]++
		right++

		// 重新组合非重复字符串
		for mp[r] > 1 {
			//移动左指针
			l := s[left]
			if _, ok := mp[l]; ok {
				mp[l]--
			}
			left++
		}

		resp = max(resp, right-left)
	}

	return resp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2. 滑动窗口
// 时间复杂度：O(2n) = O(n)，最坏的情况是 left 和 right 都遍历了一遍字符串
// 空间复杂度：O(n)
func lengthOfLongestSubstring2(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	var maxLen = 1
	var left, right, window = 0, 0, make(map[byte]bool)
	for right < n {
		var rightChar = s[right]
		// 移动左指针，直到删除map中已有的数据
		for window[rightChar] {
			fmt.Println(string(s[right]), right, window[rightChar])
			delete(window, s[left])
			left++
		}
		//if window[rightChar] {
		//	window = make(map[byte]bool)
		//	left = right
		//}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		window[rightChar] = true
		right++
	}
	return maxLen
}

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
