package main

import (
	"fmt"
	"time"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func tsstClosure() {
	a := Adder()
	ret := a(1)
	fmt.Printf("ret(1): = %d \n", ret)

	ret = a(2)
	fmt.Printf("ret(2): = %d \n", ret)

}

func tsstClosure5() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}

	time.Sleep(time.Second)
}
func main() {

	//tsstClosure()
	tsstClosure5()
}
