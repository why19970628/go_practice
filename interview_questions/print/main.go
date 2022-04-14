package main

import (
	"fmt"
	"time"
)

var (
	signal1 = make(chan struct{}, 1)
	signal2 = make(chan struct{}, 1)
)

func main() {
	//ch := make(chan int, 1)
	go printInt1()
	go printInt2()
	time.Sleep(time.Second * 1)

}

func printInt1() {
	for i := 1; i < 9; i += 2 {
		if i != 1 {
			<-signal1
		}
		fmt.Println(i)
		signal2 <- struct{}{}
	}
}

func printInt2() {
	for i := 2; i < 9; i += 2 {
		<-signal2
		fmt.Println(i)
		signal1 <- struct{}{}
	}
}
