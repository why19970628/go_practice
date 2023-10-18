package liskov_substitution

import (
	"errors"
	"fmt"
)

//继承有一些优点：
//
//1. 提高代码的重用性，子类拥有父类的方法和属性；
//2. 提高代码的可扩展性，子类可形似于父类，但异于父类，保留自我的特性；
//缺点：侵入性、不够灵活、高耦合
//1. 继承是侵入性的，只要继承就必须拥有父类的所有方法和属性，在一定程度上约束了子类，降低了代码的灵活性；
//2. 增加了耦合，当父类的常量、变量或者方法被修改了，需要考虑子类的修改，所以一旦父类有了变动，很可能会造成
//非常糟糕的结果，要重构大量的代码。

//因为继承带来的侵入性，增加了耦合性，也降低了代码灵活性，父类修改代码，子类也会受到影响，此时就需要里氏替换原则。
//
//子类必须实现父类的抽象方法，但不得重写（覆盖）父类的非抽象（已实现）方法。
//子类中可以增加自己特有的方法。
//当子类覆盖或实现父类的方法时，方法的前置条件（即方法的形参）要比父类方法的输入参数更宽松。
//当子类的方法实现父类的抽象方法时，方法的后置条件（即方法的返回值）要比父类更严格。

type IBadBird interface {
	ID() int
	Name() string
	Tweet() error
	Fly() error
}

type BadNormalBird struct {
	iID   int
	sName string
}

func NewBadNormalBird(iID int, sName string) IBadBird {
	return &BadNormalBird{iID: iID, sName: sName}
}

func (me *BadNormalBird) ID() int {
	return me.iID
}

func (me *BadNormalBird) Name() string {
	return me.sName
}

func (me *BadNormalBird) Tweet() error {
	return nil
}

func (me *BadNormalBird) Fly() error {
	fmt.Printf("BadNormalBird.Fly, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}

// BadOstrichBird.go
//不好的设计.
//BadOstrichBird通过继承BadNormalBird实现了IBadBird接口. 由于不支持Fly, 因此Fly方法抛出了错误. 额外添加了IBadBird未考虑到的Run方法. 该方法的调用要求调用方必须判断具体类型, 导致严重耦合.

type BadOstrichBird struct {
	BadNormalBird
}

func NewBadOstrichBird(id int, name string) IBadBird {
	return &BadOstrichBird{
		*(NewBadNormalBird(id, name).(*BadNormalBird)),
	}
}

func (me *BadOstrichBird) Fly() error {
	return errors.New(fmt.Sprintf("BadOstrichBird.Fly, cannot fly, id=%v, name=%v\n", me.ID(), me.Name()))
}

func (me *BadOstrichBird) Run() error {
	fmt.Printf("BadOstrichBird.Run, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
