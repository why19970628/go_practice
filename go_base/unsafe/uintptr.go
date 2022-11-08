package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

func main() {
	var w = new(W)
	fmt.Println(w)
	c := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b) + unsafe.Offsetof(w.c))
	*((*int64)(c)) = 20
	fmt.Println(w)
}
