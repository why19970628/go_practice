package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)
	sign := 1
	abs := 0
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = abs*10 + int(s[i]-'0')
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	fmt.Println(abs, sign)
	return sign * abs

}

func main() {
	//fmt.Println("123")
	//fmt.Println("-123")
	fmt.Println(myAtoi("42"))
}
