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
	"runtime"
	"syscall"
	"time"
)

type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Name  string `json:"name"`

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

	maxProces := runtime.NumCPU()
	if maxProces > 1 {
		maxProces -= 1
	}
	// Go 1.5 版本之前，默认使用的是单核心执行。从 Go 1.5 版本开始，默认执行上面语句以便让代码并发执行，最大效率地利用 CPU。
	//GOMAXPROCS 同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用。
	//此处还要注意点一点就是： Go默认执行使用的CPU核心数为系统CPU最大核心。

	// 多核比较适合那种 CPU 密集型程序，如果是 IO 密集型使用多核会增加 CPU 切换的成本。

	runtime.GOMAXPROCS(maxProces)


	router.POST("/hello", func(c *gin.Context) {
		var smsLoginParam SmsLoginParam
		//context.Request.Body 包含请求参数
		err := Decode(c.Request.Body, &smsLoginParam)
		if err != nil {
			fmt.Printf("参数解析失败~~,原因是%v\n", err.Error())
			return
		}
		smsLoginParam.Name = "WANG"
		ffmt.Puts(smsLoginParam)

		bt, err := json.Marshal(smsLoginParam)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bt))
	})

	//mp := make(map[string]interface{})





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
