package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 8, 9, 0, 0, 7, 0, 2}
	//sort.Ints(arr)
	//fmt.Println(arr)

	var newArr []int
	for i, v := range arr {
		if v == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			newArr = append(newArr, v)
		}
	}
	fmt.Println(arr)
	fmt.Println(newArr)
	arr = append(arr, newArr...)
	fmt.Println(arr)

}
