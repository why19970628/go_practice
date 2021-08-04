package main

import "fmt"

func print(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 20; i++ { //启动20个协程处理消息队列中的消息
		go print(i)
	}
	select {} // 阻塞
}
