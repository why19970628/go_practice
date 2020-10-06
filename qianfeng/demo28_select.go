package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)


	select {
	case num1 := <-ch1:
		fmt.Println("ch1 data..", num1)
	case num2, ok := <-ch2:
		if ok{
			fmt.Println("ch2 data..", num2)
		}else {
			fmt.Println("ch2 通道关闭..", num2)

		}
	case <- time.After(3 *time.Second):
		fmt.Println("case3 run .... timeout")
	//default:
	//	fmt.Println("no chan input")

	}

	fmt.Println("main over")
}