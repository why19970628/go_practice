package main

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/reverse-string/
*/
func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func TestReverseString(t *testing.T) {
	bt := []byte("123")
	reverseString(bt)
	fmt.Println(string(bt))
}
