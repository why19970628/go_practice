package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(sl []int) int {
	sort.Ints(sl)
	fmt.Println(sl)
	for i, v := range sl {
		if sl[i] == sl[i+1] {
			return v
		}
	}
	return -1

}

func findRepeatNumber2(sl []int) int {
	m := make(map[int]int)
	for i, v := range sl {
		val, ok := m[v]
		if ok && val == 1 {
			return i
		} else {
			m[v]++
		}
	}
	return -1

}

func main() {
	sl := []int{4, 1, 3, 4, 5, 6}
	fmt.Println(findRepeatNumber2(sl))
}
