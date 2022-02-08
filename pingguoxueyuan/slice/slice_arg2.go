package main

import "fmt"

func main() {
	st := make([]int, 1, 3)
	fmt.Printf("st 原始切片的地址为：%p\n", &st)
	fmt.Println("外部1:", st, len(st), cap(st))
	SliceTest(st)                             // st引用同一块“slice数组”
	fmt.Println("外部2:", st, len(st), cap(st)) // [0] 1 3

	a := st[1:]
	fmt.Println(a, len(a), cap(a))
}

// 什么意思？就是在函数传递值时，都会拷一个副本传递给函数进行操作，就算是引用类型或者指针也不例外。
// 可以看到，go语言中仅仅拷贝了st的一个副本进入SliceTest函数中进行操作，并未操作st本身。

// 在进入SliceTest函数时，s和st指向的底层数组都一样。
func SliceTest(s []int) {
	//s[0] = 5 //修改这个数组内容，所以里面改了外面能看到改动
	fmt.Printf("st 传入后参数s的地址为：%p\n", &s)
	fmt.Println("内部1:", s, len(s), cap(s))
	s = append(s, 3) //如果没有多余的空间，其实外部的是没有修改的，会创建一个新的slice
	// 如果有多余的空间，其实外部的是被修改了，只不过外部h的size还是之前的大小，
	fmt.Println("内部2:", s, len(s), cap(s))

	fmt.Printf("st 传入后参数s2的地址为：%p\n\n", &s)

}
