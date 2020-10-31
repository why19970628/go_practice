package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ffmt.v1"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}


func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	router.POST("/hello", func(c *gin.Context) {
		var smsLoginParam SmsLoginParam
		//context.Request.Body 包含请求参数
		err := Decode(c.Request.Body, &smsLoginParam)
		if err != nil {
			fmt.Printf("参数解析失败~~,原因是%v\n", err.Error())
			return
		}
		ffmt.Puts(smsLoginParam)
	})



	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit
	log.Println("Server exiting")


	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	//// kill 默认会发送 syscall.SIGTERM 信号
	//// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	//// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	//// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
	//<-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
	//log.Println("Shutdown Server ...")
	//// 创建一个5秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown: ", err)
	//}
	//log.Println("Server exiting")

	//  air https://www.jianshu.com/p/45215fe48030
}
