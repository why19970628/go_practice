package main

import (
	"fmt"
)

func main() {
	var n = 5
	if n > 10 {
		goto lable
	}
	fmt.Println("不走这里, goto继续执行")
lable:
	fmt.Println("走这里")
}

//结果输出：走这里
