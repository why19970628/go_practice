// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package main

import (
	"flag"
	"fmt"

	"pingguoxueyuan/rpc/hello/hello"
	"pingguoxueyuan/rpc/hello/internal/config"
	"pingguoxueyuan/rpc/hello/internal/server"
	"pingguoxueyuan/rpc/hello/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/hello.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewHelloServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		hello.RegisterHelloServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
