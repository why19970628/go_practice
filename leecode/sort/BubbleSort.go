package main

import "fmt"

// 冒泡排序
var array3 []int = []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}

func main() {
	count := len(array3)

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if array3[i] > array3[j] {
				swap(array3, i, j)
			}
		}
	}
	fmt.Println("res:", array3)

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
