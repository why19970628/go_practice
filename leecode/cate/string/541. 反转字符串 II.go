package main

import "fmt"

func reverseStr(s string, k int) string {
	bt := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		if (i + k) > n {
			reverse(bt[i:n])
		} else {
			reverse(bt[i : i+k])
		}

	}
	return string(bt)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
	// 输出："bacdfeg"
}
