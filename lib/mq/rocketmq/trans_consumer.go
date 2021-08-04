package main

import (
	"fmt"
	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	"strings"
	"time"
)

func main() {
	// 设置HTTP接入域名（此处以公共云生产环境为例）
	endpoint := ""
	// AccessKey 阿里云身份验证，在阿里云服务器管理控制台创建
	accessKey := ""
	// SecretKey 阿里云身份验证，在阿里云服务器管理控制台创建
	secretKey := ""
	// 所属的 Topic
	topic := ""
	// Topic所属实例ID，默认实例为空
	instanceId := ""
	// 您在控制台创建的 Consumer ID(Group ID)
	groupId := ""

	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")

	mqConsumer := client.GetConsumer(instanceId, topic, groupId, "")

	for {
		endChan := make(chan int)
		respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					// 处理业务逻辑。
					var handles []string
					fmt.Printf("Consume %d messages---->\n", len(resp.Messages))
					for _, v := range resp.Messages {
						handles = append(handles, v.ReceiptHandle)
						fmt.Printf("\tMessageID: %s, PublishTime: %d, MessageTag: %s\n"+
							"\tConsumedTimes: %d, FirstConsumeTime: %d, NextConsumeTime: %d\n"+
							"\tBody: %s\n"+
							"\tProps: %s\n",
							v.MessageId, v.PublishTime, v.MessageTag, v.ConsumedTimes,
							v.FirstConsumeTime, v.NextConsumeTime, v.MessageBody, v.Properties)

					}

					// NextConsumeTime前若不确认消息消费成功，则消息会被重复消费。
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样。
					ackerr := mqConsumer.AckMessage(handles)
					if ackerr != nil {
						// 某些消息的句柄可能超时，会导致消息消费状态确认不成功。
						fmt.Println(ackerr)
						if errAckItems, ok := ackerr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem); ok {
							for _, errAckItem := range errAckItems {
								fmt.Printf("\tErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s\n",
									errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
							}
						} else {
							fmt.Println("ack err =", ackerr)
						}
						time.Sleep(time.Duration(3) * time.Second)
					} else {
						fmt.Printf("Ack ---->\n\t%s\n", handles)
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					// Topic中没有消息可消费。
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						fmt.Println("\nNo new message, continue!")
					} else {
						fmt.Println(err)
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					fmt.Println("Timeout of consumer message ??")
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息。
		// 长轮询表示如果Topic没有消息，则客户端请求会在服务端挂起3s，3s内如果有消息可以消费则立即返回响应。
		mqConsumer.ConsumeMessage(respChan, errChan,
			3, // 一次最多消费3条（最多可设置为16条）。
			3, // 长轮询时间3秒（最多可设置为30秒）。
		)
		<-endChan
	}
}
