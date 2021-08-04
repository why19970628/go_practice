/**
 @Author: wangtong
 @Date: 2020/10/21 9:13 下午
 @Description: 钉钉Talk TODO
**/
package sd_dingding

import (
	"fmt"
	"github.com/liunian1004/dingtalk"
)

type DingTalk struct {
	*dingtalk.DingTalk
}

// 实例
func NewClient() *DingTalk {
	// 创建DingTalk实例。
	client := dingtalk.NewDingTalk(&dingtalk.Context{
		AccessToken: "",
	})
	return &DingTalk{
		DingTalk: client,
	}
}

func (d *DingTalk) SendLog(content string) error {
	err := d.SendLog(content)
	return err
}

func (d *DingTalk) SendText(content string) error {
	fmt.Println(content)
	err := d.SendText(content)
	return err
}
