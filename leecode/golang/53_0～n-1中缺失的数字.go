package main

import (
	"fmt"
	"sort"
)

func findleakNumber(sl []int) int {
	if sl[0] == 1 {
		return 0
	}
	sort.Ints(sl)
	for i, v := range sl {
		if sl[i]+1 != sl[i+1] {
			return v + 1
		}
	}
	return -1
}

func main() {
	sl := []int{7, 1, 3, 4, 5, 6}
	fmt.Println(findleakNumber(sl))
}
