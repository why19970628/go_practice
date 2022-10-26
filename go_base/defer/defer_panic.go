package main

import "fmt"

func init() {
	fmt.Println("1st init print in test package")
}

func init() {
	fmt.Println("2nd init print in test package")
}

func f(num int) {
	fmt.Println(num)
}

func main() {
	defer f(1)
	defer f(2)

	panic("123")
	defer f(3)
}
