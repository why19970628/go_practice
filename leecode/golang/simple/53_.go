package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	v := sort.SearchInts(nums, target)
	ans := 0
	for i := v; i < len(nums); i++ {
		if nums[i] == target {
			ans++
		} else {
			break
		}
	}
	return ans
}

func search2(nums []int, target int) int {
	return sort.SearchInts(nums, target+1) - sort.SearchInts(nums, target)
}

func search3(nums []int, target int) int {

	return sort.SearchInts(nums, target+1) - sort.SearchInts(nums, target)
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
	fmt.Println(search2(nums, 8))
}
