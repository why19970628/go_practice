package main

import (
	"fmt"
	"sync"
)

func run() {
	producer := make(chan int)

	wait := &sync.WaitGroup{}
	wait.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wait.Done()
			for {
				select {
				case val, ok := <-producer:
					if !ok {
						return
					}
					fmt.Println(val)
				}
			}
		}()
	}
	for i := 0; i < 10000; i++ {
		producer <- i
	}
	close(producer)
	wait.Wait()
}

func run2() {
	producer := make(chan int, 10000)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		producer <- i
	}
	close(producer)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-producer:
					if !ok {
						return
					}
					fmt.Println(val)
				}

			}

		}()
	}
	wg.Wait()
}

func main() {
	// channel队列 1生产 3消费
	// 消费完退出
	run2()
}
