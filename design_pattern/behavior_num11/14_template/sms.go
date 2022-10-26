package main

import (
	"fmt"
)

type ISMS interface {
	send(content string, phone int) error
}

type sms struct {
	ISMS
}

func (s *sms) Send(content string, phone int) error {
	return s.send(content, phone)
}

// TelecomSms 走电信通道
type TelecomSms struct {
	*sms
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}

func main() {
	tel := NewTelecomSms()
	err := tel.Send("test", 1239999)
	if err != nil {
		fmt.Println(err)
	}
}
