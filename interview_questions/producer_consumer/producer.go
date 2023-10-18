package main

import (
	"fmt"
	"time"
)

func Producer(head string, pipeline chan string) {
	for {
		pipeline <- fmt.Sprintf("%s: %v", head, time.Now().String())
	}
}

func Consumer(pipeline chan string) {
	for {
		fmt.Println(<-pipeline)
	}
}

func main() {
	channel := make(chan string)
	go Producer("dog", channel)
	go Producer("cat", channel)
	Consumer(channel)
}
