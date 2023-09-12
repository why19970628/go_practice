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

func reverseStringV2(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {

		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}
func main() {

	bt := []byte("hello")
	reverseStringV2(bt)
	fmt.Println(string(bt))
}
