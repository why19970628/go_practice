package main

import (
	"fmt"
	"sort"
)

func oddEvenSort(nums []int) {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i]%2 == 1 && nums[j]%2 == 0
	})
}

func oddEvenSortV2(nums []int) {
	slow, fast := 0, 1
	for fast < len(nums) {

		if nums[fast]%2 == 1 && nums[slow]%2 == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow]%2 == 1 {
			slow++
		}
		fast++
	}
}

func main() {
	nums := []int{2, 1, 2, 3, 4, 5}
	oddEvenSortV2(nums)
	fmt.Println(nums)
}
