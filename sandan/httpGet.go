package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ffmt.v1"
	"io/ioutil"
	"net/http"
	"time"
)

//动态图片
type ImgMeta struct {
	FileSize    ImgValue
	Format      ImgValue
	ImageHeight ImgValue
	ImageWidth  ImgValue
}
type ImgValue struct {
	Value string
}

//get 请求
func HttpGet(httpurl string) (result ImgMeta, err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", httpurl, nil)
	if err != nil {
		return ImgMeta{}, err
	}
	resp, err := client.Do(req)
	ffmt.Puts(resp.Body)

	if err != nil {
		ffmt.Puts(err)
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return ImgMeta{}, err2
	}
	err3 := json.Unmarshal([]byte(string(body)), &result)

	return result, err3

}

func main() {
	url := "https://img.ipcfun.com/uploads/ishoulu/pic/2013/05/9215193abd60b5ff099795216.jpg"
	meta, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(meta)
}
