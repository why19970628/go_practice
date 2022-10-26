package router

import (
	"encoding/json"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go_practice/frame/gin_test/controllers/opentracing"
	"go_practice/frame/gin_test/middware"
	"io"
	"net/http"
	"time"
)

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

func InitRouter() *gin.Engine {

	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.Use(middware.LoggerHandler(), middware.OpenTracing())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	testGroup := router.Group("/test")
	{
		testGroup.GET("/rpc", opentracing.Rpc)
		testGroup.GET("/sms", opentracing.Sms)
		testGroup.GET("/panic", opentracing.Panic)
		testGroup.GET("/timeout", timeout.New(
			timeout.WithTimeout(1000*time.Microsecond),
			timeout.WithHandler(opentracing.EmptySuccessResponse),
		))
	}
	return router
}
