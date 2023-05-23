package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		currByte := pairs[s[i]]
		fmt.Println(stack)
		// 如果是)]}
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != currByte {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
func main() {
	fmt.Println(isValid("{[()]}"))
}
