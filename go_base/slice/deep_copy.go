package main

import "fmt"

type HandlersChain struct {
	A string
}

func main() {
	two()
	three()
	four()
}

func one() {
	finalSize := 10
	mergedHandlers := make([]int, finalSize)
	fmt.Println(mergedHandlers)

	slice1 := []int{1, 2, 3, 4, 5}
	copy(mergedHandlers, slice1)
	fmt.Println(mergedHandlers)
	fmt.Println("----")

	slice2 := []int{5, 4, 3}
	fmt.Println(len(slice1), mergedHandlers[len(slice1):])
	i := copy(mergedHandlers[len(slice1):], slice2)
	fmt.Println(i, mergedHandlers)
}

func two() {
	s := []string{"A", "B", "C"}

	t := s[1:2]
	fmt.Println(&s[0] == &t[0])
	fmt.Printf("%p, %p \n", s, t)

	fmt.Println("==========")
	u := append(s[:1], s[2:]...)
	fmt.Println(&s[0] == &u[0])
	u[0] = "6666"
	fmt.Printf("s=%v u= %v 地址%p, %p \n", s, u, s, u)

}

func three() {
	fmt.Println("=========three")
	arr := [3]int{1, 2, 3}
	arr2 := arr
	arr3 := &arr
	fmt.Printf("原数组：%v \n", arr)         // 123
	fmt.Printf("赋值方式复制的数组：%v \n", arr2)  // 123
	fmt.Printf("引用方式复制的数组：%v \n", *arr3) // 123
	arr[1] = 1000
	fmt.Printf("改变后原数组：%v \n", arr)         // 1-1000-3
	fmt.Printf("改变后赋值方式复制的数组：%v \n", arr2)  // 123
	fmt.Printf("改变后引用方式复制的数组：%v \n", *arr3) // 1-1000-3
}

func four() {
	fmt.Println("=========four")
	arr := []int{1, 2, 3}
	fmt.Printf("1--- arr= %v 地址%p, \n", arr, arr)
	temp(arr)
	fmt.Printf("3--- arr= %v 地址%p, \n", arr, arr)

}

func temp(arr []int) {
	arr = append(arr, 4)
	fmt.Printf("2--- arr= %v 地址%p, \n", arr, arr)
}
