package main

import (
	"encoding/json"
	"fmt"
)

func check(str []byte, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}
}

func main() {
	a := 3
	p := &a
	var ch = make(chan int, 1)
	var c complex64 = 1i + 3
	var f = func() { fmt.Println("hello") }

	str, err := json.Marshal(p)
	check(str, err)

	check(json.Marshal(a))
	check(json.Marshal(ch))
	check(json.Marshal(c))
	check(json.Marshal(f))
}

// 结果只有指针会被强转成相应的数据并输出
