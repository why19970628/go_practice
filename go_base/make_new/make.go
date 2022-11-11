package main

import "fmt"

// 二者异同
// 1.二者都是内存的分配（堆上），
// 2.但是make只用于slice、map以及channel的初始化（非零值）；
// 而new用于类型的内存分配，并且内存置为零。
// 3.make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
func testMake() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 66
	a = append(a, 11, 20, 12)
	fmt.Println(a) // [66 0 0 0 0 11 20 12]
	// make也是用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

}

func make2() {
	var i *int
	*i = 10
	fmt.Println(*i) // panic
	// 从这个提示中可以看出，对于引用类型的变量，我们不光要声明它，还要为它分配内容空间，否则我们的值放在哪里去呢？
	// 这就是上面错误提示的原因。

}

func make3() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)
	// 这就是new，它返回的永远是类型的指针，指向分配类型的内存地址。

}

func main() {
	testMake()
	//make2()
	make3()
}

func make4() {
	//直接 n3:=[]int {1,2} 和 make 的存放地址不一样
	//n3==0xc000014070
	//若 n2:=make([]int,0)
	//n2==0x1181f88
	//所以说还是有不一样的。
	//使用make的好处是可以指定len和cap，make(type,len,cap),合适的len和cap可以提升性能。
}
