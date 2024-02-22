package main

import "fmt"

func firstUniqueChar(str string) string {
	mp := make(map[byte]int, len(str))
	for i := 0; i < len(str); i++ {
		mp[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if mp[str[i]] == 1 {
			return string(str[i])
		}
	}
	return " "
}

func main() {
	fmt.Println(firstUniqueChar("leecode"))
}
