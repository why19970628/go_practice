package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)

		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1 data..", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2 data..", num2)
		} else {
			fmt.Println("ch2 通道关闭..", num2)

		}
		//default:
		//	fmt.Println("no chan input")

	}

	fmt.Println("main over")
}
