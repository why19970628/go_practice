package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"sync"
	"time"
)

func main() {
	hystrix.ConfigureCommand("get_baidu", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i:=0; i<100; i++ {
		go TestHystix(i, &wg)
		//time.Sleep(1*time.Second)
	}
	wg.Wait()
	time.Sleep(2 * time.Second) // 调用Go方法就是起了一个goroutine，这里要sleep一下，不然看不到效果
}


func TestHystix(i int, wg *sync.WaitGroup)  {
	// 根据自身业务需求封装到http client调用处
	hystrix.Go("get_baidu", func() error {
		// 调用关联服务
		res, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Println("get error")
			return err
		}
		//time.Sleep(time.Second * 2) // 请模拟求超时
		fmt.Println("请求成功：",i, res.Status)
		return nil
	},
		// 失败重试，降级等具体操作
		func(err error) error {
			fmt.Println("get an error, handle it:", err)
			return nil
		})
	wg.Done()
}