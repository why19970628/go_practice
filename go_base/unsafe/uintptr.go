package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

type Programmer struct {
	name     string
	language string
}

func f2() {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)
	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	language := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*language = "golang"
	fmt.Println(p)
}

func f1() {
	var w = new(W)
	fmt.Println(w)
	c := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b) + unsafe.Offsetof(w.c))
	*((*int64)(c)) = 20
	fmt.Println(w)
}

func main() {
	f2()
}
