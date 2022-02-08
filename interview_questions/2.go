package main

import "fmt"

//func search2(arr []int, target int) int {
//	if len(arr) < 1 {
//		return -1
//	}
//	mid := len(arr) / 2
//	left := 0
//	right := len(arr) - 1
//	for i := 0; i < len(arr); i++ {
//
//	}
//	//for _, v := range arr {
//	//	if arr[mid] > target {
//	//		left =
//	//	}else {
//	//
//	//	}
//	//}
//
//}
func main() {

	arr := [8]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	s := arr[1:2]
	fmt.Println(s)
	s[0] = 1111
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("----")

	arr2 := []int{1, 2, 3, 4, 5, 6}
	s2 := arr2[1:5]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
