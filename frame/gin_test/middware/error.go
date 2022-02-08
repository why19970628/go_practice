package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		RequestURI := c.Request.Method + "  " + c.Request.Host + c.Request.RequestURI
		defer func() {
			if err := recover(); err != nil {
				event := fmt.Sprintf(
					"*时间:%v\n"+
						//"*RESPONSE STATUS: %d\n"+
						"*描述: 接口500\n"+
						"*URI: %v \n"+
						"*error:= %s\n"+
						"*RequestIP:= %v\n"+
						"",
					t,
					RequestURI,
					err,
					c.ClientIP(),
				)
				fmt.Println(event)
			}
		}()
		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		body, _ := ioutil.ReadAll(c.Request.Body)

		if status != 500 || latency > time.Second {
			if status >= 500 {
				fmt.Println(c.FullPath(), c.Request.RequestURI)
				fmt.Printf("PATH: %v | USE TIME: %v | RESPONSE STATUS: %d\n", RequestURI, latency, c.Writer.Status())
			} else {
				fmt.Printf("*接口超时 PATH: %v\n"+
					"*耗时 %v\n"+
					"body %v\n", RequestURI, latency, string(body))
			}
		}
	}
}
