package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
			if a == 9 {
				fmt.Println("close2")
				return
			}
		}
	}()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("ok")
}
