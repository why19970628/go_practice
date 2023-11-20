package main

import "fmt"

// abc ac
func maxHuiWen(str string) int {
	// 记录数量
	mp := make(map[byte]int)
	l, r := 0, 0
	n := len(str)
	resp := 0
	for r < n {
		mp[str[r]]++
		for mp[str[r]] > 1 {
			resp = max(resp, r-l)
			mp[str[l]]--
			// 移动左指针
			l++
		}
		// 移动右指针
		r++
	}
	return resp

}
func main() {
	fmt.Println(maxHuiWen("abcac"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
