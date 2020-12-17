package main

import (
	"fmt"
	"reflect"
)

type People struct {
}

func (p *People) Student() {
	fmt.Println("i am student!!!")
}
func main() {
	p := People{}
	v := reflect.ValueOf(&p)
	f := v.MethodByName("Student")
	fmt.Println(f)
	f.Call([]reflect.Value{})
}
