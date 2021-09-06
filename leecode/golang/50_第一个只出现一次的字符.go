package main

import (
	"fmt"
)

func FirstNotRepeatingChar(str string) int {
	charMap := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		char := str[i]
		//fmt.Println(reflect.TypeOf(char))
		if val, ok := charMap[char]; ok {
			charMap[char] = val + 1
		} else {
			charMap[char] = 1
		}
	}

	for i := 0; i < len(str); i++ {
		if charMap[str[i]] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(FirstNotRepeatingChar("abcdea"))
}
