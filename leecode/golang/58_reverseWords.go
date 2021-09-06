package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	tmp := strings.Fields(strings.Trim(s, " "))
	reverse(tmp)
	fmt.Println(tmp)
	join := strings.Join(tmp, " ")
	return strings.Trim(join, "")
}
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	words := "abc"
	s := reverseWords(words)
	fmt.Println(s)
}
