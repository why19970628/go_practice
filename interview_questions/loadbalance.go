package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Instance struct {
	host string
	port int
}

type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}

type RoundRobinBalance struct {
	curIndex int
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	fmt.Println(p.curIndex)
	p.curIndex = (p.curIndex + 1) % lens
	//fmt.Println(p.curIndex)
	return
}

func main() {
	var insts []*Instance
	for i := 0; i < 100; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := NewInstance(host, 8080)
		insts = append(insts, one)
		fmt.Println(one)
	}

	var d Balancer
	d2 := &RoundRobinBalance{}
	d = d2
	for i := 0; i < 1000; i++ {
		//res, err := d.DoBalance(insts, "random")
		//if err != nil {
		//	fmt.Println(err)
		//}
		res, err := d.DoBalance(insts, "random")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("res: ", res)
	}
}
