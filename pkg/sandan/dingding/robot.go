package sd_dingding

/**
 * @Author: long
 * @Description:
 * @File:  DingTalk
 * @Version: 1.0.0
 * @Date: 2020-05-27 19:35
 */

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ffmt.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	url string = "https://oapi.dingtalk.com/robot/send?access_token=289e7f82b1ffa602998069531854fe977bc06f3c61b356c149b5145e211778b4"
)

/*  desc robot send
 *  doc  https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq
 */
func RobotSend(code int, sendurl string, data map[string]interface{}) (result map[string]interface{}, err error) {
	var info map[string]interface{}
	if len(sendurl) > 0 {
		url = sendurl
	}
	switch code {
	case 1:
		info = map[string]interface{}{
			"msgtype": "text",
			"text": map[string]interface{}{
				"content": "文字:" + data["content"].(string),
			},
			"at": map[string]interface{}{
				"atMobiles": data["mobiles"].([]string),
				"isAtAll":   data["isAtAll"].(bool),
			}}
	case 2:
		info = map[string]interface{}{
			"msgtype": "link",
			"link": map[string]interface{}{
				"text":       "文字:" + data["text"].(string),
				"title":      data["title"].(string),
				"picUrl":     data["picUrl"].(string),
				"messageUrl": data["messageUrl"].(string),
			},
		}
	case 3:
		info = map[string]interface{}{
			"msgtype": "markdown",
			"markdown": map[string]interface{}{
				"title": "文字：" + data["title"].(string),
				"text":  data["text"].(string),
			},
			"at":      data["mobiles"].([]string),
			"isAtAll": data["isAtAll"].(bool),
		}
	case 4: //整体跳转
		info = map[string]interface{}{
			"msgtype": "actionCard",
			"actionCard": map[string]interface{}{
				"title":          "文字：" + data["title"].(string),
				"text":           data["text"].(string),
				"btnOrientation": data["btnOrientation"].(string),
				"singleTitle":    data["singleTitle"].(string),
				"singleURL":      data["singleURL"].(string),
			},
		}
	case 5: //独立跳转
		/*"btns": [
		    {
		        "title": "内容不错",
		        "actionURL": "https://www.dingtalk.com/"
		    },
		    {
		        "title": "不感兴趣",
		        "actionURL": "https://www.dingtalk.com/"
		    }
		]*/
		info = map[string]interface{}{
			"msgtype": "actionCard",
			"actionCard": map[string]interface{}{
				"title":          "文字" + data["title"].(string),
				"text":           data["text"].(string),
				"btnOrientation": data["btnOrientation"].(string),
				"btns":           data["btns"].([]map[string]interface{}),
			},
		}
	case 6:
		/*"links": [
		    {
		        "title": "时代的火车向前开",
		        "messageURL": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
		        "picURL": "https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"
		    },
		    {
		        "title": "时代的火车向前开2",
		        "messageURL": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
		        "picURL": "https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"
		    }
		]*/
		info = map[string]interface{}{
			"msgtype": "feedCard",
			"feedCard": map[string]interface{}{
				"links": data["links"].([]map[string]interface{}),
			},
		}
	}
	contenttype := "json"
	headers := map[string]interface{}{
		"Accept":       "application/" + contenttype,
		"Content-Type": "application/" + contenttype + ";charset=utf-8",
	}
	result, err = HttpPost(url, info, headers)
	ffmt.Puts(err)
	ffmt.Puts(result)
	return
}

//客服消息预警
func RobotNotifySysMessage(data map[string]interface{}, dingUrl string) (result map[string]interface{}, err error) {
	info := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": data["title"].(string),
			"text":  data["text"].(string),
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   false,
		},
	}
	contenttype := "json"
	headers := map[string]interface{}{
		"Accept":       "application/" + contenttype,
		"Content-Type": "application/" + contenttype + ";charset=utf-8",
	}
	result, err = HttpPost(dingUrl, info, headers)
	fmt.Printf("result: %v, err:%v \n", result, err)
	return
}

//post 请求
func HttpPost(httpurl string, data map[string]interface{}, headers map[string]interface{}) (result map[string]interface{}, err error) {

	client := &http.Client{}
	info, err := json.Marshal(data)
	req, err := http.NewRequest("POST", httpurl, strings.NewReader(string(info)))
	if err != nil {
		return nil, err
	}
	for k, va := range headers {
		req.Header.Set(k, va.(string))
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}
	result = map[string]interface{}{}
	err3 := json.Unmarshal([]byte(string(body)), &result)

	return result, err3

}
