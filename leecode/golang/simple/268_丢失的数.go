package main

import (
	"fmt"
	"sort"
)

// 1.hash
// 2.loop
func missNumber(num []int) int {
	sort.Ints(num)
	for i, v := range num {
		if i+1 != v {
			return i + 1
		}
	}
	return -1
}

func missNumber2(num []int) int {
	for i, v := range num {
		if i < len(num) {
			if v+1 != num[i+1] {
				return v + 1
			}
		}
		if i == len(num) {
			return v
		}
	}
	return -1
}
func main() {
	num := []int{2, 5, 4, 1}
	fmt.Println(missNumber(num))
	fmt.Println(missNumber2(num))

}
