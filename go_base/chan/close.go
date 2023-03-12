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

func closeCh() {
	// chan不需要显示关闭(close)，只要没有goroutine持有channel，相关资源会自动释放。
	// chan最好是发送端调用close，此时读取端会收到一个空消息，如下：
	closeCh := make(chan string)
	go func() {
		for {
			select {
			case i := <-closeCh:
				if i == "" {
					fmt.Println("通道被close")
					return
				} else {
					fmt.Println("收到消息:", i)
				}
			default:

			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			closeCh <- "123"
		}
		close(closeCh)
	}()
	time.Sleep(time.Second * 1)
}

func main() {
	//closeCh()

	signalChannel := make(chan struct{})
	go Done(signalChannel, 1)
	go Done(signalChannel, 2)
	go Done(signalChannel, 3)
	time.Sleep(time.Second * 1)
	fmt.Println("close signalChannel")
	close(signalChannel)

	time.Sleep(time.Second * 2)
	// 阻塞当前goroutine
	//select {}
}
