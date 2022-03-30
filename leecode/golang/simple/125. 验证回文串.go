package main

import (
	"fmt"
	"strings"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
func checkhuiwei() bool {
	s := "A man, a plan, a canal: Panama"
	var sgood = ""
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood = sgood + string(s[i])
		}
	}
	fmt.Println(sgood)

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	fmt.Println(sgood)
	fmt.Println(n)

	for i := 0; i < n; i++ {
		if sgood[i] != sgood[n-i-1] {
			return false
		}
	}
	return true

}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	fmt.Println(checkhuiwei())
}
