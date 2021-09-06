package main

import (
	"fmt"
	"gopkg.in/ffmt.v1"
)

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// stu是临时变量,地址值不变, 所以&stu会把相同的地址值一直带到最后一次循环
	for _, stu := range stus {

		fmt.Printf("%+v\n", &stu)
		fmt.Printf("%+v\n", stu)
		tmp := stu
		m[stu.Name] = &tmp
	}

	ffmt.Puts(m)

	for k, v := range m {
		println(k, "=>", v.Name)
	}

}
