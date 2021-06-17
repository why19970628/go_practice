package main

import (
	"fmt"
	"sync"
	"time"
)

func TestMap() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	m := make(map[int]int)

	l := sync.Mutex{}
	for j := 0; j < 100; j++ {
		go func(t int) {
			l.Lock()
			m[t] = t
			l.Unlock()
		}(j)
		wg.Done()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(m)

	for j := 0; j < 100; j++ {
		go func(j int) {
			if _, ok := m[j]; ok {
				fmt.Println("get data", j)
			} else {
				fmt.Println("no data")
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}

func main() {
	TestMap()
}
