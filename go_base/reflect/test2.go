package main

import (
	"fmt"
	"reflect"
)

type Animal1 struct{}

func (a *Animal1) Do() {
	fmt.Println("test")
}

func main() {
	a := Animal1{}
	v := reflect.ValueOf(&a)
	f := v.MethodByName("Do")
	f.Call([]reflect.Value{})
}
