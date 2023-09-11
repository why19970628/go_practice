package main

import (
	"fmt"
	"testing"
)

func MaxSubHuiWen(s string) string {
	resp := ""
	for i := 0; i < len(s); i++ {
		s1 := expand(i, i, s)
		s2 := expand(i, i+1, s)
		resp = max1(s1, s2, resp)
	}

	return resp
}

func expand(left, right int, s string) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func max1(a, b, c string) string {
	resp := a
	if len(b) > len(resp) {
		resp = b
	}
	if len(c) > len(resp) {
		resp = c
	}
	return resp

}

func TestMaxSubHuiWen(t *testing.T) {
	fmt.Println(MaxSubHuiWen("12321"))
}
