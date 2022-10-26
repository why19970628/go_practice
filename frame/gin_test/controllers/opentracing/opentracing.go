package opentracing

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	rpc_http "go_practice/frame/gin_test/libraries/http"
)

func Rpc(c *gin.Context) {
	postData := make(map[string]interface{})
	postData["query"] = [1]string{"猕猴桃"}

	//logCfg := config.GetConfigToJson("log", "log")
	//logId := c.Writer.Header().Get(logCfg["query_field"].(string))
	logId := "123"
	sendUrl := "https://www.baidu.com"

	ret := rpc_http.HttpSend(c, "GET", sendUrl, logId, postData)
	ret = rpc_http.HttpSend(c, "GET", sendUrl, logId, postData)

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "success",
		"data":   ret,
	})
	c.Done()
}

func Panic(c *gin.Context) {
	panic("test err")

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "success",
		"data":   make(map[string]interface{}),
	})
	c.Done()
}

func Long(c *gin.Context) {
	time.Sleep(time.Second * 3)
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "success",
		"data":   make(map[string]interface{}),
	})
	c.Done()
}

func EmptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "sucess")
	return
}

type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

type s struct {
	a SmsLoginParam
	b string
}

func Sms(c *gin.Context) {

	a := s{
		a: SmsLoginParam{
			Name: "21312",
			Code: "131",
		},
		b: "123",
	}
	time.Sleep(time.Second)
	//panic("123")

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "success",
		"data":   a,
	})
	c.Done()
}
