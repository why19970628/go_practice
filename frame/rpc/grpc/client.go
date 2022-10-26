package main

import (
	"context"
	"fmt"
	pb "go_practice/frame/rpc/grpc/protos"
	"google.golang.org/grpc"
)

func main() {

	// 客户端连接服务器
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
	}
	defer conn.Close()

	// 获得grpc句柄
	c := pb.NewHelloServerClient(conn)

	// 远程单调用 SayHi 接口
	r1, err := c.SayHi(
		context.Background(),
		&pb.HelloRequest{
			Name: "Kitty",
		},
	)
	if err != nil {
		fmt.Println("Can not get SayHi:", err)
		return
	}
	fmt.Println("SayHi 响应：", r1)

	// 远程单调用 GetMsg 接口
	r2, err := c.GetMsg(
		context.Background(),
		&pb.HelloRequest{
			Name: "Kitty",
		},
	)
	if err != nil {
		fmt.Println("Can not get GetMsg:", err)
		return
	}
	fmt.Println("GetMsg 响应：", r2)

}
