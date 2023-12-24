package main

import "fmt"

func sortArray(nums []int) []int {
	sortArray_quickSort(nums, 0, len(nums)-1)
	return nums
}

func sortArray_quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := quickSortPartion(nums, start, end)
	sortArray_quickSort(nums, start, pivot-1)
	sortArray_quickSort(nums, pivot+1, end)
}

func quickSortPartion(nums []int, start, end int) int {
	l, r := start, start
	pivot := nums[end]
	for r < end {
		if nums[r] < pivot {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
		r++
	}
	nums[l], nums[end] = nums[end], nums[l]
	return l
}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
