package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, taskChan chan int) {
	defer wg.Done()
	for {
		select {
		case i, ok := <-taskChan:
			if !ok {
				return
			}
			//time.Sleep(time.Second * 1)
			fmt.Println("task:", i)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout...")
			return
		case <-ctx.Done():
			fmt.Println("收到信号，退出worker")
			return
		}
	}

}
func master(workerNums, taskNum int) {
	ctx, cancel := context.WithCancel(context.Background())
	taskChan := make(chan int, taskNum)
	wg := sync.WaitGroup{}
	for i := 0; i < workerNums; i++ {
		wg.Add(i)
		go worker(ctx, &wg, taskChan)
	}

	for i := 0; i < taskNum; i++ {
		taskChan <- i
	}

	close(taskChan)
	wg.Wait()
	cancel()
	fmt.Println("Master finished")
}

func main() {
	workerNums := 3
	taskNum := 1000
	master(workerNums, taskNum)
}
