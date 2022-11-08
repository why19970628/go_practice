package main

import (
	"fmt"
	"time"
)

func main() {

	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}() //每次将变量 v 的拷贝传进函数
	}
	time.Sleep(time.Second)
	fmt.Println("=========")
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v) //每次将变量 v 的拷贝传进函数
	}
	time.Sleep(time.Second)

	fmt.Println("=========2")

	t4()
	time.Sleep(time.Second)
}

// 3，延迟调用
func b() {
	x, y := 1, 2

	defer func(a int) {
		fmt.Printf("x:%d,y:%d\n", a, y) // y 为闭包引用
	}(x) // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)
	// 101 102
	//x:1,y:102
}

type myType int

func (v *myType) show4() {
	fmt.Printf("t4 pointer addr: %v addr in pointer: %v & value in pointer: %v \n", &v, v, *v)
}

func t4() {
	values := []myType{1, 2, 3, 4, 5}
	for _, val := range values {
		fmt.Printf("t4 for loop addr & value:%v %v\n", &val, val)
		go val.show4()
	}
}
