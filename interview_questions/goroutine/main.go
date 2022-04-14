package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	names := []string{"lily", "yoyo", "cersei", "rose", "annei"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
		time.Sleep(time.Second)
	}
	runtime.GOMAXPROCS(1)
	runtime.Gosched()
}
