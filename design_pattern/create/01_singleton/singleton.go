package singleton

import (
	"fmt"
	"sync"
)

// golang单例模式
// 1、定义：单例对象的类必须保证只有一个实例存在，全局有唯一接口访问。

// 应用场景：一个无状态的类使用单例模式节省内存资源。

type Singleton interface {
	foo()
}

type singleton struct {
	singleton1
}

func (s *singleton) foo() {
	fmt.Println("foo")
	s.singleton1.c = "c"
	s.singleton1.Name = "Name"
	fmt.Println(s)
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})
	return instance
}
