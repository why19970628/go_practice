package singleton

import (
	"fmt"
	"sync"
)

type si interface {
	t()
}

type singleton1 struct {
	Name string
	c    string
}

func (s *singleton1) t() {
	fmt.Println("tttt")
}

// 饿汉模式
var b *singleton1

func init() {
	b = &singleton1{}
}

func GetInstance1() *singleton1 {
	return b
}

// 懒汉模式
var once2 = &sync.Once{}

func GetInstance2() *singleton1 {
	if b == nil {
		once2.Do(func() {
			b = &singleton1{}
		})
	}
	return b
}
