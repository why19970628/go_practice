package main

import (
	"fmt"
	"sync"
)

var a = make([]int, 40)

func main() {
	lock := new(sync.Mutex)
	for i := 0; i < 40; i++ {
		go func(i int) {
			lock.Lock()
			a[i]++
			lock.Unlock()
		}(i)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
}
