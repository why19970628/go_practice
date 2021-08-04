package main

import (
	"github.com/haisum/rpcexample"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	//注册Arith 对象作为一个服务
	arith := new(rpcexample.Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalf("Format of service Arith isn't correct. %s", err)
	}
	//将Rpc绑定到HTTP协议上
	rpc.HandleHTTP()
	//开始监听端口 1234的消息
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	//启动http服务，处理连接请求l
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
