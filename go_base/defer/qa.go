package main

import "fmt"

// 容易掉坑之，函数变量修改
func f3() (res int) {
	defer func() {
		res++
	}()
	return 0
}

// 容易掉坑之，参数复制
func f4() (res int) {
	defer func(res int) {
		res++
	}(res)
	return 0
}

func main() {
	fmt.Println(f3()) // 1
	fmt.Println(f4()) // 0
}
