package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	Name string
}

func (c *Client) Do() {
	fmt.Println("client do....")
}

type LoadBalancer struct {
	client []*Client
	size   int32
}

func NewLoadBalancer(size int32) *LoadBalancer {
	loadBalancer := &LoadBalancer{client: make([]*Client, size), size: size}
	loadBalancer.client = append(loadBalancer.client, &Client{})
	loadBalancer.size = loadBalancer.size + 1
	return loadBalancer
}

func (m *LoadBalancer) getClient() *Client {
	rand.Seed(time.Now().Unix())
	x := rand.Int31n(100)
	fmt.Println(x)
	fmt.Println(x % m.size)
	return m.client[x%m.size]
}

func main() {
	lib := NewLoadBalancer(4)
	lib.getClient().Do()
	fmt.Println(lib.size, lib)
}
