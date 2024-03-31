package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addStrings1(num1 string, num2 string) string {
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	t := 0
	result := strings.Builder{}
	for n1 >= 0 || n2 >= 0 || t > 0 {
		var a, b int
		if n1 >= 0 {
			a = int(num1[n1] - '0')
			n1--
		}
		if n2 >= 0 {
			b = int(num2[n2] - '0')
			n2--
		}

		sum := a + b + t

		t = sum / 10
		c := sum % 10
		result.WriteString(strconv.Itoa(c))
	}
	s := result.String()
	rs := []byte(s)
	l, r := 0, len(rs)-1
	for l < r {
		rs[l], rs[r] = rs[r], rs[l]
		l++
		r--
	}
	return string(rs)
}

func main() {
	num1 := "12345678901234567890"
	num2 := "98765432101234567890"

	sum := addStrings(num1, num2)
	fmt.Println("Sum:", sum)
}
