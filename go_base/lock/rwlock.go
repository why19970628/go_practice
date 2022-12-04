package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	count int
	wg    sync.WaitGroup
	rw    sync.RWMutex
)

func main() {
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		go write(i)
	}

	wg.Wait()
}

func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)

	v := count

	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
	rw.Unlock()
}
