package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "aaa"
	fmt.Println(unsafe.Pointer(&a))                     // unsafe.Pointer(&a) 方法可以得到变量a的指针地址。
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a)) // 把字符串a转成 StringHeader底层结构形式 的实际内容
	fmt.Println(fmt.Sprintf("%+v", ssh))                //  {Data:17601749 Len:3}
	b := *(*[]byte)(unsafe.Pointer(&ssh))               // 字符串a转成底层结构形式。
	fmt.Printf("%v", b)
	v2()
}

func v2() {
	b := "bbb"
	strData := *(*reflect.StringHeader)(unsafe.Pointer(&b))
	res := *(*[]byte)(unsafe.Pointer(&strData))
	fmt.Println(res, string(res))
}
