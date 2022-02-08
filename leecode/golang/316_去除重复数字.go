package main

import (
	"fmt"
	"sort"
)

func solution() {
	s := "bcabc"
	m := make(map[string]int)
	m2 := make(map[int]string)

	for index, value := range s {
		fmt.Println(value)
		if _, ok := m[string(value)]; !ok {
			m[string(value)] = index
		}

	}
	fmt.Println(m)

	slice := make([]int, 0)
	for i, v := range m {
		m2[v] = i
		slice = append(slice, v)
	}
	fmt.Println(m2)

	sort.Ints(slice)
	res := ""
	for _, v2 := range slice {
		res += m2[v2]
	}
	fmt.Println(res)

}

func main() {
	solution()
}
