package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"io"
	"net/http"
	"pingguoxueyuan/frame/gin_test/controllers/opentracing"
	"pingguoxueyuan/frame/gin_test/middware"
)

type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

func InitRouter() *gin.Engine {

	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.Use(middware.OpenTracing())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	testGroup := router.Group("/test")
	{
		testGroup.GET("/rpc", opentracing.Rpc)
		testGroup.GET("/panic", opentracing.Panic)
	}
	return router
}
