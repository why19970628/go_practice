package main

import "fmt"

func search(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] > target {
			right = mid - 1
		} else if array[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

type localInt int
type Map map[string]interface{}

type i = localInt
type localMap = Map

func main() {
	//array := []int{1, 2, 3, 4, 5}
	//val := 1

	array2 := []int{1, 2, 4, 5}
	val2 := 3
	//res := search(array, val)
	res2 := search(array2, val2)

	//fmt.Println(res)
	fmt.Println(res2)
}
