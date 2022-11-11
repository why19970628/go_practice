package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Printf("%p\n%v\n", fp, *fp) // 地址一样

	*fp = *fp * 3

	fmt.Println(i)
}
