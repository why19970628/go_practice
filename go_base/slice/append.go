package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Printf("%v %p\n", x, &x) //output: 0xc420012240, 0xc420012240, 0xc420012248
	y := x[:2]
	fmt.Printf("222:%v %p\n", y, &y) //output: 0xc420012240, 0xc420012240, 0xc420012248
	y = append(y, 50)                // [1 2 50]
	fmt.Printf("333 y : %v %p\n", y, &y)
	fmt.Printf("333 x: %v %p\n", x, &x)

	fmt.Println("-------------------")
	y = append(y, 60)            // 超出容量,公用底层的切片会重新分配一个地址 x y 此时独立
	fmt.Printf("%v %p\n", y, &y) //output: [1 2 50 60]
	y[0] = 20
	fmt.Printf("%v %p\n", y, &y)        //output:  [20 2 50 60]
	fmt.Printf("555 x: %v %p\n", x, &x) // [1 2 50] 0xc00000c060
}
