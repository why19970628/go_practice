package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chanLimitRate()
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
