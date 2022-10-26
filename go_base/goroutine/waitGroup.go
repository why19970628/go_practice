package main

import (
	"fmt"
	"runtime"
	"sync"
)

func max() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	max()
}

//func main() {
//	//runtime.GOMAXPROCS(1)
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i int) {
//			fmt.Println(i)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
