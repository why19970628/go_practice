package main

import "fmt"

func main() {
	var s1 = []int{1, 2}    // 初始化一个切片
	var s2 = make([]int, 2) // 初始化一个空的切片，cap为2
	copy(s2, s1)            // 将s1拷贝给s2
	s3 := s1                // 将s1拷贝给s2

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s2[0] = 99         // 改变s2[0]
	s3[0] = 1001       // 改变s2[0]
	fmt.Println(s1[0]) // 打印 1 而不是 99

	fmt.Println("---")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
