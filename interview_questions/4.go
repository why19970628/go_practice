package main

import (
	"fmt"
	"sort"
)

func solution4(num int) {
	arr := []int{1, 1}
	for i := 1; i < num; i++ {
		v := arr[i-1] + arr[i]
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func main() {
	//num := 10
	//solution4(num)

	arr := []int{89, 3, 8, 67, 78, 44, 63, 9}
	//n := []int{}
	//for index, v := range a {
	//	s := strconv.Itoa(v)
	//	fmt.Println(s, string(s[0]))
	//	str := string(s[0])
	//	i, _ := strconv.Atoi(str)
	//	//n = append(n, a[index])
	//}

	m := make(map[int]int)

	newarr := []int{}
	for _, v := range arr {
		if v < 10 {
			m[v*10] = v
			v = v * 10
		}
		newarr = append(newarr, v)
	}
	fmt.Println(m)
	fmt.Println(newarr)

	sort.Ints(newarr)
	fmt.Println(newarr)

	newarr2 := []int{}
	for i := len(newarr) - 1; i >= 0; i-- {
		v := newarr[i]
		if v2, ok := m[v]; ok {
			v = v2
		}
		newarr2 = append(newarr2, v)

	}
	fmt.Println(newarr2)

}
