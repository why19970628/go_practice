package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

//轮询负载均衡
type WeightRoundRobinBalance struct {
	curIndex int
	rss      []string
}

func NewWeightRoundRobinBalance() *WeightRoundRobinBalance {
	return &WeightRoundRobinBalance{}
}

func (w *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) <= 1 {
		return errors.New("权重必须大于0")
	}
	addr := params[0]
	weight, err := strconv.Atoi(params[1])
	fmt.Println(addr, weight)
	if err != nil {
		return err
	}
	if weight == 0 {
		weight = 1
	}
	for i := 0; i < weight; i++ {
		w.rss = append(w.rss, addr)
	}
	return nil
}
func (w *WeightRoundRobinBalance) Next() string {
	if len(w.rss) <= 0 {
		return ""
	}
	w.curIndex = rand.Intn(len(w.rss))
	return w.rss[w.curIndex]
}

func (r *WeightRoundRobinBalance) Get() (string, error) {
	return r.Next(), nil
}

func main() {
	r := NewWeightRoundRobinBalance()
	for i := 0; i < 10; i++ {
		r.Add(strconv.Itoa(i), strconv.Itoa(rand.Intn(10)))
	}
	m := make(map[string]int)
	for i := 0; i < 1000; i++ {
		res, _ := r.Get()
		//fmt.Println(res)
		m[res]++
	}
	fmt.Println(m)
}
