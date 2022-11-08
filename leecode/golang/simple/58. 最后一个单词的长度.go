package main

import "fmt"

//func lengthOfLastWord(s string) int {
//	index := len(s) - 1
//	for s[index] == ' ' {
//		index--
//	}
//	var ans int
//	for index > 0 && s[index] != ' ' {
//		fmt.Println(string(s[index]))
//		ans++
//		index--
//	}
//
//	return ans
//}

func lengthOfLastWord(s string) int {
	var ans int
	// 逆序
	lastIndex := len(s) - 1
	for s[lastIndex] == ' ' {
		lastIndex--
	}
	for s[lastIndex] != ' ' {
		ans++
		lastIndex--
	}

	return ans
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World   "))
}
