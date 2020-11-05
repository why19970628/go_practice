package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载自定义的配置文件
	InitDb()
	r := gin.New()
	r.Run()
}
