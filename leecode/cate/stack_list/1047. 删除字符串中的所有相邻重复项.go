package main

import "fmt"

func removeDuplicates(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	stack := make([]byte, 0)

	stack = append(stack, s[:1]...)
	for i := 1; i < n; i++ {
		currByte := s[i]

		var StackLastByte byte
		if len(stack) > 0 {
			StackLastByte = stack[len(stack)-1]
		}
		if currByte != StackLastByte {
			stack = append(stack, currByte)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
