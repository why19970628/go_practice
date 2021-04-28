package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	//"gin-api/libraries/logging"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	opentracingLog "github.com/opentracing/opentracing-go/log"
)

func HttpSend(c *gin.Context, method, url, logId string, data map[string]interface{}) map[string]interface{} {
	//ctx := c.Request.Context()
	var (
		//logFormat = logging.LogHeaderFromContext(ctx)
		err error
		ret = make(map[string]interface{})
		req *http.Request
	)

	//if logFormat == nil {
	//	logFormat = &logging.LogHeader{}
	//}

	//logFormat.LogId = logId

	client := &http.Client{}

	//请求数据
	byteDates, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(byteDates)

	//url
	url = url + "?logid=" + logId

	//构建req
	req, err = http.NewRequest(method, url, reader)
	if err != nil {
		panic(err)
	}

	//设置请求header
	req.Header.Add("content-type", "application/json")

	tracer, _ := c.Get("Tracer")
	parentSpanContext, _ := c.Get("ParentSpanContext")

	span := opentracing.StartSpan(
		"httpDo",
		opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		ext.SpanKindRPCClient,
	)
	defer span.Finish()

	injectErr := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	if injectErr != nil {
		span.LogFields(opentracingLog.String("inject-error", err.Error()))
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ret["code"] = resp.StatusCode
	ret["msg"] = "success"
	ret["data"] = make(map[string]interface{})

	if resp.StatusCode != http.StatusOK {
		ret["msg"] = "http code:" + strconv.Itoa(resp.StatusCode)
	}

	if b != nil {
		res, err := simplejson.NewJson(b)
		if err != nil {
			ret["data"] = string(b)
		} else {
			ret["data"] = res
		}
	}

	span.SetTag("url", url)
	span.SetTag("code", resp.StatusCode)
	span.SetTag("msg", ret["msg"])
	span.SetTag("data", ret["data"])

	return ret
}
