package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < threads-1; i++ {

		// go 1.13卡死所有的p ，主goroutine 无法退出
		go func(x *int) {
			for {
				*x++
			}
		}(&x)
	}

	for i := 0; i < threads; i++ {
		// go 1.13卡死所有的p ，主goroutine 无法退出
		go func() {
			for {
				x++
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(x)
}
