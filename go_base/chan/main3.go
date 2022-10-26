package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 1)
	go func() {
		for {
			v, ok := <-intChan
			if !ok {
				break
			} else {
				fmt.Println(v)
			}
		}
	}()
	intChan <- 1
	close(intChan)
	time.Sleep(time.Second * 1)
}
