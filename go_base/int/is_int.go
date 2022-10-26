package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1.23
	b := ""
	if a == float64(int64(a)) {
		fmt.Println("yay")
		b = string(int(a))
	} else {
		fmt.Println("you fail")
		b = strconv.FormatFloat(a, 'E', -1, 64) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	}
	fmt.Println(b)
}
