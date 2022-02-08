package adapter

//Target 是适配的目标接口
type Target interface {
	Request() string
	Request2() string
}

//Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
	SpecificRequest2() string
}

//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//AdapteeImpl 是被适配的目标类
type adapteeImpl struct{}

//SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest2() string {
	return "adaptee method2"
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

//Request 实现Target接口
func (a *adapter) Request2() string {
	return a.SpecificRequest2()
}
