package main

import "fmt"

func s(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

var res int = -1

func s2(arr []int, target, left, right int) int {
	mid := (left + right) / 2
	if arr[mid] > target {
		s2(arr, target, left, mid-1)
	} else if arr[mid] < target {
		s2(arr, target, mid+1, right)
	} else {
		fmt.Println(mid)
		res = mid
		return res
	}
	return res
}

func main() {
	arr := []int{1, 3, 5, 6, 7, 9, 13}
	target := 9
	fmt.Println(s(arr, target))
	fmt.Println(s2(arr, target, 0, len(arr)-1))
}
