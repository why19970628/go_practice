package main

import (
	"errors"
	"fmt"
	"strconv"
)

//轮询负载均衡
type RoundRobinBalance struct {
	curIndex int
	rss      []string
}

func NewRoundRobinBalance() *RoundRobinBalance {
	return &RoundRobinBalance{}
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}

	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	curElement := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % len(r.rss)
	return curElement
}

func (r *RoundRobinBalance) Get() (string, error) {
	return r.Next(), nil
}

func main() {
	r := NewRoundRobinBalance()
	for i := 0; i < 10; i++ {
		r.Add(strconv.Itoa(i))
	}
	for i := 0; i < 100; i++ {
		res, _ := r.Get()
		fmt.Println(res)
	}
}
