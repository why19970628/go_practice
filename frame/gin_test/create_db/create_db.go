package main

import (
	"github.com/gin-gonic/gin"
)

// sql := "SELECT * FROM exchange_match_pays where TIMESTAMPDIFF(SECOND, created_at, NOW()) > 15 AND is_deleted= 0 AND deleted_at is  null"
func main() {
	// 加载自定义的配置文件
	InitDb()
	r := gin.New()
	r.Run()
}
