package main

import "fmt"

/*
1.简单选择排序
基本思想:在L[i+1]到L[n]中选择关键字最小的元素与L[i]进行交换。
*/
func selectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		var minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		Swap(arr, i, minIndex)
	}
	return arr

}

func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	fmt.Println(selectSort([]int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}))
}
