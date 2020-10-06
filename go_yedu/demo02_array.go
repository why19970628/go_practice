package main

import "fmt"

var x1 = []int{2:5, 6, 0:7}

func main() {
	x := []int{123}
	x, x[0] = nil, 456

	fmt.Println(x) //[]

	fmt.Println(x1) // [7 0 5 6]
}
