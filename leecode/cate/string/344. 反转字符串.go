package main

import "fmt"

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left <= right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		right--
		left++
	}
}

func main() {

	bt := []byte("hello")
	reverseString(bt)
	fmt.Println(string(bt))
}
