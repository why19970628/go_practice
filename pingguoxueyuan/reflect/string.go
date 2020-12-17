package main

import (
	"fmt"
	"reflect"
)

type Animal struct {

}

func (a *Animal)Eat()  {
	fmt.Println("Eat")
}

func main()  {
	animal := Animal{}
	v := reflect.ValueOf(&animal)
	//t := reflect.TypeOf(&animal)
	//fmt.Println(v.Kind())
	f := v.MethodByName("Eat")
	f.Call([]reflect.Value{})
}