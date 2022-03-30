package liskov_substitution

import "fmt"

//_可以理解为: _
//所有引用父类的地方
//必须能透明地使用其子类对象
//子类对象能够替换父类对象
//而保持程序功能不变

//里氏替换原则的优点:
//（1）约束继承泛滥，是开闭原则的一种体现
//（2）加强程序的健壮性，同时变更时可以做到非常好的兼容性
//
type IGoodBird interface {
	ID() int
	Name() string

	Tweet() error
}

type IFlyableBird interface {
	IGoodBird

	Fly() error
}

type IRunnableBird interface {
	IGoodBird

	Run() error
}

// GoodNormalBird.go
// GoodNormalBird提供对IGoodBird的基础实现
type GoodNormalBird struct {
	iID   int
	sName string
}

func NewGoodNormalBird(id int, name string) *GoodNormalBird {
	return &GoodNormalBird{
		id,
		name,
	}
}

func (me *GoodNormalBird) ID() int {
	return me.iID
}

func (me *GoodNormalBird) Name() string {
	return me.sName
}

func (me *GoodNormalBird) Tweet() error {
	fmt.Printf("GoodNormalBird.Tweet, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}

// GoodOstrichBird通过聚合GoodNormalBird实现IGoodBird接口, 通过提供Run方法实现IRunnableBird子接口

type GoodOstrichBird struct {
	GoodNormalBird
}

func NewGoodOstrichBird(id int, name string) IGoodBird {
	return &GoodOstrichBird{
		*NewGoodNormalBird(id, name),
	}
}

func (me *GoodOstrichBird) Run() error {
	fmt.Printf("GoodOstrichBird.Run, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
