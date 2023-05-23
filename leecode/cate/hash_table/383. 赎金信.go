package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[rune]int)
	for _, v := range ransomNote {
		mp[v]++
	}
	for _, v := range magazine {
		if value, ok := mp[v]; ok {
			mp[v] = value - 1
		}
	}
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))

}
