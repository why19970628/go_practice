package main

import "fmt"

func main() {
	t2()
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
		fmt.Println(stu)
		//a := &stu
		//fmt.Printf("aaaa, %p\n", a)
		//m[stu.Name] = &stu

		//fmt.Printf("%+v\n", &stu)
		//fmt.Printf("%+v\n", stu)
		//tmp := stu
		//m[stu.Name] = &tmp
	}

	//ffmt.Puts(m)

	for k, v := range m {
		println(k, "=>", v.Name)
	}

}

func pase_student2() {
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// stu是临时变量,地址值不变, 所以&stu会把相同的地址值一直带到最后一次循环
	for _, stu := range stus {
		fmt.Println(stu)
		//a := stu
		fmt.Printf("aaaa, %p\n", stu)
		//m[stu.Name] = stu

		//fmt.Printf("%+v\n", &stu)
		//fmt.Printf("%+v\n", stu)
		//tmp := stu
		m[stu.Name] = stu
	}
	fmt.Println(m)

	//ffmt.Puts(m)

	for k, v := range m {
		println(k, "=>", v.Name)
	}

}

func t2() {
	var nums = []int64{1, 2, 3, 4}
	var t []*int64
	for _, num := range nums {
		fmt.Println(num)
		//fmt.Printf("aaaa, %p", &num,)
		//a := num
		t = append(t, &num)
	}
	//fmt.Println(t)
	for _, v := range t {
		fmt.Println(v, *v)
	}
}
