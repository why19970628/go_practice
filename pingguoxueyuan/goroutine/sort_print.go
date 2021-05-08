package main

import (
	"sync"
	"time"
)

var ch chan struct{}
var ch2 chan struct{}

func p1() {
	time.Sleep(time.Millisecond * 3)
	println("p1")
	ch <- struct{}{}
}

func p2() {
	<-ch
	println("p2")
	ch2 <- struct{}{}
}

func p3() {
	time.Sleep(time.Millisecond * 10)
	println("p3")
	<-ch2
}

func do(fns ...func()) {
	wg := sync.WaitGroup{}
	for _, fn := range fns {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
	wg.Wait()
}
func main() {
	ch = make(chan struct{})
	ch2 = make(chan struct{})
	do(p1, p2, p3)
}
