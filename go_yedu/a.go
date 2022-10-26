package main

import (
	"fmt"
	"net/url"
	"time"
)

//1其中 urlvalues 的定义：
type Values map[string][]string

type Query struct {
	url.Values
}

func main() {
	q := Query{}
	q.Values = make(map[string][]string)
	q.Values["name"] = []string{"moonfdd1"}
	fmt.Println(q.Get("name"))
	v2()

}
func v2() {
	ch := make(chan string)
	//endChan := make(chan int)

	go func() {
		fmt.Println("go func start....")
		for i := 0; i < 100; i++ {
			ch <- "result"
			time.Sleep(time.Second * 1)
		}
	}()

	select {
	//第一个case里阻塞的时间只有比第二个case阻塞的时间长的时候, 才能执行第二个case
	case res := <-ch:
		fmt.Println(res)
		//endChan <- 1
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
		//endChan <- 1
	}
	//<-endChan
}
