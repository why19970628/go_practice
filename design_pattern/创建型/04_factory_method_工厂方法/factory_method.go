package factory_method

// 模式特点：定义一个用于创建对象的接口，让子类决定实例化哪一个类。这使得一个类的实例化延迟到其子类。解耦多个产品之间的业务逻辑。

//  工厂函数在增加一个新的类型的时候，像上面的例子里，只要增加一个新类型的工厂类，里面实现

type OperationI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

type Operation struct {
	a float64
	b float64
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

type AddOperation struct {
	Operation
}

func (this *AddOperation) GetResult() float64 {
	return this.a + this.b
}

type SubOperation struct {
	Operation
}

func (this *SubOperation) GetResult() float64 {
	return this.a - this.b
}

type MulOperation struct {
	Operation
}

func (this *MulOperation) GetResult() float64 {
	return this.a * this.b
}

type DivOperation struct {
	Operation
}

func (this *DivOperation) GetResult() float64 {
	return this.a / this.b
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {
}

func (this *AddFactory) CreateOperation() OperationI {
	return &(AddOperation{})
}

type SubFactory struct {
}

func (this *SubFactory) CreateOperation() OperationI {
	return &(SubOperation{})
}

type MulFactory struct {
}

func (this *MulFactory) CreateOperation() OperationI {
	return &(MulOperation{})
}

type DivFactory struct {
}

func (this *DivFactory) CreateOperation() OperationI {
	return &(DivOperation{})
}
