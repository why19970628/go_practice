package main

import (
	"fmt"
	"math"
	"strconv"
)

func nextPalindrome(num int) int {
	for i := num + 1; i < math.MaxInt64; i++ {
		if isPalindrome(i) {
			return i
		}
	}
	return -1
}
func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	l, r := 0, len(str)-1
	for l <= r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true

}
func main() {
	fmt.Println(nextPalindrome(123))
}
