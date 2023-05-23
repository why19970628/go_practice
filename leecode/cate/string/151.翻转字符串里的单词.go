package main

import "fmt"

func reverseWords(s string) string {
	bt := []byte(s)
	n := len(bt)
	reverseStr1(bt, 0, n-1)
	fmt.Println(string(bt))
	return string(bt)
}

func reverseStr1(b []byte, left, right int) {
	for left <= right {
		tmp := b[right]
		b[right] = b[left]
		b[left] = tmp
		left++
		right--
	}
}

func main() {
	reverseWords("ab cd")
}
