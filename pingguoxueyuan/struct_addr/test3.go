package main

import "fmt"

type People struct {
	Name string
	Age int
}


type Person struct {
	Firstname string
	Lastname string
	Age uint8
}
// 值方法
func (p Person) show() {
	fmt.Println(p.Firstname)
}
// 指针方法
func (p *Person) show2() {
	fmt.Println(p.Firstname)
}

// 值方法
func (p Person) setFirstName(name string) {
	p.Firstname = name
}
// 指针方法
func (p *Person) setFirstName2(name string) {
	p.Firstname = name
}



// 当调用值方法setFirstName后，输出的还是原来的值sun，
//而调用指针方法 setFirstNam2后，则输出的是新值。主要原因就是值方法func (p Person)在传递总结体的时候，用的只是原来结构体的一个副本，
//做的任何修改也只是对副本的修改，而打印的还是原来的结构体，两者互不影响。
//而指针方法，传递的则是指向结构体指针的值副本，指针值一样(X012242R424)，指定的都是底层的数据结构，所以才会出现这种情况。
func main() {
	p := Person{"sun", "xingfang", 30}
	//不一致的情况
	p.show() // sun 修改前
	p.setFirstName("tom")   // 值方法
	p.show() // sun, 未变化
	p.show() // sun 修改前
	fmt.Println("-------")
	p.setFirstName2("tom")  // 指针方法
	p.show() // tom 修改后的tom

	fmt.Println("2-------")

	var p2 People
	p2.Name = "wang"
	p2.Age = 12
	fmt.Println(p2) // {wang 12}

	p3 := p2
	p3.Name = "hua"
	p3.Age = 14
	fmt.Println(p3) // {hua 14}
	fmt.Println("2-------")


	peo := new(People)
	fmt.Println(peo)              //&{ 0}
	fmt.Println(peo == nil)       //false

	peo.Name = "derek"
	peo.Age = 22
	fmt.Println(peo)              //&{derek 22}

	peo2 := peo
	fmt.Println(peo2)            //&{derek 22}

	peo2.Name = "Jack"
	fmt.Println(peo, peo2)       //&{Jack 22} &{Jack 22}

}


