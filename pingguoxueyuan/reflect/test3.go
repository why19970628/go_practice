package main

import (
	"fmt"
	"reflect"
)


type user struct {
	Name string
	Age int
}
func (u *user) ShowName() {
	fmt.Println(u.Name)
}

func (u *user) AddAge(addNum int) {
	fmt.Println("age add result:", u.Age + addNum)
}

func main()  {
	var num int64 = 100

	// 设置值：指针传递
	ptrValue := reflect.ValueOf(&num)
	fmt.Println(ptrValue)
	newValue := ptrValue.Elem()                         // Elem()用于获取原始值的反射对象
	fmt.Println(newValue)
	fmt.Println("type：", newValue.Type())				// int64
	fmt.Println(" can set：", newValue.CanSet())		// true
	newValue.SetInt(200)

	// 获取值：值传递
	rValue := reflect.ValueOf(num)
	fmt.Println(rValue.Int())					// 方式一：200
	fmt.Println(rValue.Interface().(int64))		// 方式二：200


	u := &user{"list", 20}
	v := reflect.ValueOf(u)
	f := v.MethodByName("ShowName")
	f.Call(nil)

	// 调用有参方法
	methodV2 := v.MethodByName("AddAge")
	args := []reflect.Value{reflect.ValueOf(30)}	//
	methodV2.Call(args)
}
