package main

import (
	"fmt"
	"time"
)

//Functional Option Pattern
//属于是创建者模型
//我们创建实例有很多方法，其本质都是创建结构体
//但缺点是，直接去构造结构体代码的可读性并不高，也并不是良好的接口

// 同样也是option的方式进行扩展
type Server struct {
	host    string
	port    int
	timeout time.Duration
}

func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	svr := &Server{
		host: addr,
	}
	for _, option := range options {
		option(svr)
	}
	return svr, nil
}

func timeout(d time.Duration) func(*Server) {
	return func(server *Server) {
		server.timeout = d
	}
}

func port(port int) func(*Server) {
	return func(server *Server) {
		server.port = port
	}
}

func main() {
	server, err := NewServer(
		"127.0.0.1",
		timeout(1e9),
		port(8080))
	if err != nil {
		panic(err)
	}
	fmt.Println(server)
}
