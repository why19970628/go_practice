package main

import "fmt"

func main() {
	str := "cbababsac"
	new := ""
	for i := 0; i < len(str)-1; i++ {
		for a := i + 1; a < len(str)-1; a++ {
			new = str[i:a]
			if !isnotReverseStr(new) {
				fmt.Println(new)
			}
		}
	}
	fmt.Println(isnotReverseStr(str))
}

func isnotReverseStr(s string) bool {
	var flag bool
	if len(s) <= 1 {
		return true
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[len(s)-1-i] {
			flag = true
			break
		}
	}
	return flag
}
