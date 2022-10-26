//demo_23.go
package main

import "fmt"

func main() {
	fmt.Println("begin")

	for i := 1; i <= 10; i++ {
		if i == 6 {
			// 改变函数内代码执行顺序，不能跨函数使用。
			goto END
		}
		fmt.Println("i =", i)
	}

END:
	fmt.Println("end")
}
