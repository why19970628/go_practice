package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		if value, ok := m[v]; ok {
			m[v] = value + 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range t {
		fmt.Println(string(v))
		if value, ok := m[v]; ok {
			m[v] = value - 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
