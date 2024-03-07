package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}

func main() {
	var i8 int8 // 1个字节
	fmt.Println(unsafe.Sizeof(i8))

	var i int // 8个字节
	fmt.Println(unsafe.Sizeof(i))

	var i64 int64 // 8个字节
	fmt.Println(unsafe.Sizeof(i64))

	fmt.Println(unsafe.Sizeof(Args{})) // 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 8 4+2 = 6
}
