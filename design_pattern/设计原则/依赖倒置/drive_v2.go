package 依赖倒置

import "fmt"

//Person中使用接口Drive来定义交通工具，出行依赖的是交通工具接口，而不是具体的交通工具。
//这样，可以自由选择交通工具，只要交通工具实现该接口，并将其传递给Person类中的Drive接口。
//好处：可扩展性好，以后添加其他交通工具不影响代码实现，即抽象不依赖于细节。

type Drive interface {
	drive()
}

type Bike struct{}

func (b *Bike) drive() {
	fmt.Println("drive by Bike!")
}

type Car struct{}

func (c *Car) drive() {
	fmt.Println("drive by Car!")
}

// Person类
type Person struct {
	drive Drive
}

func NewPerson() *Person {
	return &Person{
		drive: new(Bike),
	}
}
func (p *Person) DriveTool() {
	p.drive.drive()
}
