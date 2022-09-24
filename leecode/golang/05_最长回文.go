package main

import "fmt"

// https://blog.csdn.net/apple2banana/article/details/124389038
func main() {
	str := "cbababsac"
	fmt.Println(isHuiwen(str))
	//new := ""
	//for i := 0; i < len(str)-1; i++ {
	//	for a := i + 1; a < len(str)-1; a++ {
	//		new = str[i:a]
	//		if !isnotReverseStr(new) {
	//			fmt.Println(new)
	//		}
	//	}
	//}
	//fmt.Println(isnotReverseStr(str))
	f2()
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

func f2() {
	s := "cbababsac"
	max := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isHuiwen(s[i:j]) {
				res := s[i:j]
				if len(res) > len(max) {
					max = res
				}
			}

		}
	}
	fmt.Println(max)
}

func isHuiwen(s string) bool {
	if len(s) <= 1 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
