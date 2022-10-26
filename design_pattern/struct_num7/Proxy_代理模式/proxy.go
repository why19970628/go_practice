package proxy

import "fmt"

type Proxy1 interface {
	Say()
}

type proxyCtx struct {
	Name string
}

func (p *proxyCtx) Say() {
	fmt.Println(p.Name, "proxyCtx")
}

// ---------------
type now1 struct {
	Name string
	proxyCtx
}

func newNow1(name string) *now1 {
	return &now1{Name: name}
}

func (p *now1) Say() {
	fmt.Println(p.Name)
}

// ------------------------------------------------------------ //

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
