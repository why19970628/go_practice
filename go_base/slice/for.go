package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for i, val := range arr {
		arr[i] = val + 1
	}
	fmt.Println(arr)
}
