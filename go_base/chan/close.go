package main

import (
	"fmt"
	"time"
)

func Done(ch chan struct{}, count int) {
	for {
		select {
		case <-ch:
			fmt.Println(count)
			return
		}
	}
}

func main() {
	signalChannel := make(chan struct{})
	go Done(signalChannel, 1)
	go Done(signalChannel, 2)
	go Done(signalChannel, 3)
	time.Sleep(time.Second * 3)
	fmt.Println("close signalChannel")
	close(signalChannel)

	time.Sleep(time.Second * 2)
	// 阻塞当前goroutine
	select {}
}
