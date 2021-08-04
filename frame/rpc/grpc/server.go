package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "pingguoxueyuan/frame/rpc/grpc/protos"
)

// 对象要和proto内定义的服务一致
type server struct {
}

// 实现rpc的 SayHi接口
func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{
		Message: "Hi " + in.Name,
	}, nil
}

// 实现rpc的 GetMsg接口
func (s *server) GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error) {
	return &pb.HelloMessage{
		Msg: "Server msg is coming...",
	}, nil
}

func main() {

	// 监听网络
	ln, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("网络异常：", err)
		return
	}

	// 创建grpc句柄
	srv := grpc.NewServer()

	// 将server结构体注册到grpc服务中
	pb.RegisterHelloServerServer(srv, &server{})

	// 监听服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("监听异常：", err)
		return
	}

}
