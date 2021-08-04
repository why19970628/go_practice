package main

import (
	"github.com/haisum/rpcexample"
	"log"
	"net/rpc"
)

func main() {
	//与远程过程调用服务器建立连接
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//构造参数对象
	args := &rpcexample.Args{
		A: 2,
		B: 3,
	}
	//this will store returned result
	var result rpcexample.Result
	//调用远程过程Arith.Multiply，参数存放在args中，result中存放结果
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//输出result中的结果
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
