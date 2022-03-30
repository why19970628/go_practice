package 依赖倒置

// 依赖倒置原则
//依赖倒置原则（Dependence Inversion Principle, DIP）指设计代码结构时，高层模块不应该依赖底层模块，二者都应该依赖其抽象。
//除此之外 抽象不应该依赖细节，细节应该依赖抽象。
//常见的依赖注入方式有: 方法参数注入, 构造器参数注入, setter方法注入

// 好处：可扩展性好，以后添加其他交通工具不影响代码实现，即抽象不依赖于细节。

// 具体内容：针对对接口编程，依赖于抽象而不依赖于具体

type IUser interface {
	ID() int
	Name() string
	Study(ICourse)
}

type GoodUser struct {
	iID   int
	sName string
}

func NewGoodUser(id int, name string) IUser {
	return &GoodUser{
		iID:   id,
		sName: name,
	}
}

func (me *GoodUser) ID() int {
	return me.iID
}

func (me *GoodUser) Name() string {
	return me.sName
}

func (me *GoodUser) Study(course ICourse) {
	course.SetUser(me)
	course.Study()
}
