package _func

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	pb "go_practice/test/func/proto"
	"testing"
	"time"
)

func TestRedisSetProto(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Network:   "",
		Addr:      "",
		Dialer:    nil,
		OnConnect: nil,
		Username:  "",
		Password:  "",
		DB:        0,
	})

	req := &pb.MyMessage{
		Name: "testName",
		Age:  12,
	}
	var err error
	bt, _ := proto.Marshal(req)
	key := "go_base:proto:msg:1"
	err = c.Set(context.Background(), key, bt, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}

	btResp, err := c.Get(context.Background(), key).Bytes()
	if err != nil {
		panic(err)
	}

	// 反序列化消息
	retrievedMsg := &pb.MyMessage{}
	err = proto.Unmarshal(btResp, retrievedMsg)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrievedMsg::", retrievedMsg)
}
