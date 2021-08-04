package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	hystrix.ConfigureCommand("default", hystrix.CommandConfig{
		Timeout:                3000, // 单次请求 超时时间
		MaxConcurrentRequests:  1000, // 最大并发量
		SleepWindow:            5000, // 熔断后多久去尝试服务是否可用
		RequestVolumeThreshold: 1,    // 验证熔断的 请求数量, 10秒内采样
		ErrorPercentThreshold:  1,    // 验证熔断的 错误百分比
	})

	//开启一个http监控服务
	//可以使用hystrix-dashboard面板查看具体情况
	//https://github.com/mlabouardy/hystrix-dashboard-docker
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("", "8888"), hystrixStreamHandler)
		log.Fatal(err)
	}()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go Do(i, &wg)
	}
	wg.Wait()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func Do(params int, wg *sync.WaitGroup) {
	err := hystrix.Do("defalut", func() error {
		fmt.Println(params)
		//time.Sleep(1 * time.Second)
		//log.Println("sleep 2 second")
		return nil
	}, nil)

	if err != nil {
		//加入自动降级处理，如获取缓存数据等
		switch err {
		case hystrix.ErrCircuitOpen:
			fmt.Println("circuit error:" + err.Error())
		case hystrix.ErrMaxConcurrency:
			fmt.Println("circuit error:" + err.Error())
		case hystrix.ErrTimeout:
			fmt.Println("circuit time error:" + err.Error())
		default:
			fmt.Println("circuit error:" + err.Error())
		}

	}
	wg.Done()
}
