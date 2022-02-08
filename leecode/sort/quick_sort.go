package main

import "testing"

var array []int = []int{3, 5, 1, 2, 7, 8}

// 选择排序
func Test(t *testing.T) {
	count := len(array)
	var newArr []int
	for i := 0; i < count; i++ {
		index := selectionSort(array)
		newArr = append(newArr, array[index])
		array = append(array[:index], array[index+1:]...)
	}
	t.Log("res:", newArr)

}

func selectionSort(array []int) int {
	var (
		min       = array[0]
		index int = 0
	)
	for i, v := range array {
		if v <= min {
			min = v
			index = i
		}
	}
	return index
}
