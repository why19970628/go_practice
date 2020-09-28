package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArray(a [5]int) int {
	var b int
	for i := 0; i < len(a); i++ {
		b = b + a[i]
	}
	return b
}

func rondomArray()  {
	rand.Seed(time.Now().Unix())
	var b[5]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	fmt.Println(sumArray(b))
}
// 找出数组中两个元素之和为8的下标
func ArrayTwoSum()  {

	var b[5]int = [5]int{2,6,4,5,7}
	for i := 0; i < len(b); i++ {
		for a:= i+1; i<len(b); a++{

			if b[i] + b[a] ==8{
				fmt.Println(i, b[i] , a, b[a])
			}
		}
	}
	fmt.Println(sumArray(b))
}

func main() {
	var a[5]int = [5]int{100,200,300,400,500}
	fmt.Println(sumArray(a))
	rondomArray()
	ArrayTwoSum()
}
