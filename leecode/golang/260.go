package main

import "fmt"

func singleNumber(nums []int) []int {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	println(m)
	arr := []int{}
	for k, _ := range m {
		arr = append(arr, k)
	}
	return arr
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}
