package simple_factory

import "fmt"

// 专门定义一个类来负责创建其它类的实例，实例一般有共同的父类
// 模式特点：工厂类根据条件产生不同功能的运算类对象，客户端不需要知道具体的运算类。
//程序实例：四则运算计算器，根据用户的输入产生相应的运算类，用这个运算类处理具体的运算。

// 简单工厂模式在增加一个新的类型时，除了要增加新的类和方法之外，还需要修改工厂函数，在工厂函数中添加case，
// 这一点违背了对修改关闭这个原则（开放-封闭原则），所以需要工厂模式。
type Animal interface {
	Say()
}

func NewAnimalSimpleFactory(name string) Animal {
	if name == "tiger" {
		return &Tiger{}
	} else if name == "Cat" {
		return &Cat{}
	}
	return nil
}

type Tiger struct{}

func (t *Tiger) Say() {
	fmt.Println("tigger Say")
}

type Cat struct{}

func (t *Cat) Say() {
	fmt.Println("cat Say")
}
