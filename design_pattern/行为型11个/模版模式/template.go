package main

import "fmt"

// 例：我们有两个文档处理程序，一个处理本地数据，一个处理网络数据。这两个类的功能和执行步骤高度类似，但是，获取数据的方式不同。
//在这两个类中势必存在大量结构相同的代码。
//现在，我们使用模版模式来重构这两个类。

// 1.定义父类，在父类中定义实现的具体函数和一个等待子类实例化的“抽象函数”

type DocSuper struct {
	GetContent func() string
}

func (d DocSuper) DoOperate() {
	fmt.Println("对这个文档做了一些处理,文档是:", d.GetContent())
}

// 我们把这两个类，提取出一个父类，DocSuper。这个类中有两个函数，一个具体的函数，DoOperate，用来做实际的操作。
//一个抽象函数，等待子类实现，用来获取不同文档类型的内容。其实，我们只是等待子类实现这个抽象函数，好为父类的DoOperate来提供数据:)。

//另外，我们最最好用的go语言，是没有抽象类，抽象函数之类的概念的，所以我使用了一个指向一个函数的指针（func() string）来模拟实现抽象函数。
//如果我们的子类需要“实例化”这个“抽象函数”（实质是子类给这个父类函数指针赋值），就必须满足我们的指针约束。

//大家仔细看代码就能发现，我们的子类中可以包含父类的结构，虽然go语言不能继承，但是通过这种方式，我们也能过一过“继承”的瘾。
//当我们使用NewNetDoc函数实例化的时候，在这个函数中，会把子类的函数指针赋值给父类的函数指针。
//此时，子类中包含了父类的结构，故而我们能调用父类的函数来完成功能。

type LocalDoc struct {
	DocSuper
}

func NewLocalDoc() *LocalDoc {
	c := new(LocalDoc)
	c.DocSuper.GetContent = c.GetContent
	return c
}

func (e *LocalDoc) GetContent() string {
	return "this is a LocalDoc."
}

type NetDoc struct {
	DocSuper
}

func NewNetDoc() *NetDoc {
	c := new(NetDoc)
	c.DocSuper.GetContent = c.GetContent
	return c
}

func (c *NetDoc) GetContent() string {
	return "this is net doc."
}

// 总结：通过模版模式，我们可以把子类做为父类的模版，提取出公共的结构到父类，共享父类的代码。
//这样能消除代码结构重复的坏味。并且，简化了子类的功能，使之职责单一的为“父类”提供数据。

func main() {
	netDoc := NewNetDoc()
	lcDoc := NewLocalDoc()

	netDoc.DoOperate()
	lcDoc.DoOperate()
}
