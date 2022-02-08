package main

import "fmt"

func findSameNumber(nums []int) {
	m := make(map[int]int)

	for i, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = i
		} else {
			fmt.Println(v)
		}
	}
}

func main() {
	nums := []int{1, 3, 4, 2, 2}

	findSameNumber(nums)
}
