package main

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	rss []string

}
func (r*RandomBalance) Add(params ...string) error{
	if(len(params)==0){
		return errors.New("至少需要1个参数")
	}
	addr:=params[0]
	r.rss = append(r.rss,addr)
	return nil
}
func (r*RandomBalance) Next() (string,error){
	if(len(r.rss)==0){
		return "" ,errors.New("不存在参数")
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex],nil
}

