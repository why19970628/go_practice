package main

import (
	"fmt"
	"unicode"
)

var (
	age       int
	money     float32
	isMarried bool
	name      string
)

func isHan() {
	count := 0
	s := "hello, 沙河小王子"
	for _, v := range s { //rune utf-8
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func main() {
	age = 20
	money = 200.0
	isMarried = false
	name = "Villero"
	fmt.Printf("年龄：%d, %v", age, age)
	fmt.Println()
	fmt.Printf("资产：%f", money)
	fmt.Println()
	fmt.Printf("婚：%t", isMarried)
	fmt.Println()
	fmt.Printf("名：%s\n", name)

	traversalString()
	isHan()
}
