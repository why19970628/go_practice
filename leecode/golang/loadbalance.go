package main

import (
	"fmt"
	"math/rand"
)

type LoadBalance interface {
}

type Server struct {
	Host string
	Port int
	Id   int
	//主机是否在线
	Online bool
}

type LoadBalanceRandom struct {
	servers []*Server
}

// 实例化 随机均衡负载
func NewLoadBalanceRandom(servers []*Server) *LoadBalanceRandom {
	newBalance := &LoadBalanceRandom{
		servers: servers,
	}
	return newBalance
}

//选择一个后端Server
func (r *LoadBalanceRandom) Next() *Server {
	if len(r.servers) == 0 {
		return nil
	}
	curIndex := rand.Intn(len(r.servers))
	return r.servers[curIndex]
}

func (r *LoadBalanceRandom) Get(key string) (*Server, error) {
	return r.Next(), nil
}

func main() {
	count := make(map[string]int32, 4)
	servers := make([]*Server, 0)
	servers = append(servers, &Server{Host: "1", Id: 0, Online: true})
	servers = append(servers, &Server{Host: "2", Id: 1, Online: true})
	servers = append(servers, &Server{Host: "3", Id: 2, Online: true})
	servers = append(servers, &Server{Host: "4", Id: 3, Online: true})
	lb := NewLoadBalanceRandom(servers)
	// 创建4个Server，随机选择100000次。查看4台机器 被选中次数
	for i := 0; i < 100000; i++ {
		c := lb.Next()
		count[c.Host]++
	}
	fmt.Println(count)
}
