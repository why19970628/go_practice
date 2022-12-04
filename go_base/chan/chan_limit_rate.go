package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//chanLimitRate()
	chanLimitRate2()
}

// 利用channel缓冲区的特性，控制并发数量
func chanLimitRate() {
	worker := 2
	ch := make(chan struct{}, worker)
	wg := sync.WaitGroup{}
	totalCount := 1000
	wg.Add(totalCount)
	for i := 0; i < totalCount; i++ {
		i := i
		ch <- struct{}{}

		go func() {
			fmt.Println("i=", i)
			time.Sleep(time.Second * 1)
			<-ch
			wg.Done()
		}()

	}
	wg.Wait()
}

// 退出
func chanLimitRate2() {
	worker := 2
	stop := make(chan bool)
	for i := 0; i < worker; i++ {
		go func() {
			for {
				select {
				case <-stop:
					fmt.Println("监控退出，停止了...")
					return
				default:
					fmt.Println("goroutine监控中...")
					time.Sleep(2 * time.Second)
				}
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	for i := 0; i < worker; i++ {
		stop <- true
	}
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(1 * time.Second)
}
