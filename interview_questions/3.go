package main

import "fmt"

func search2(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	fmt.Println(left, right)
	for left <= right {
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

func search3(arr []int, target, left, right int) int {
	mid := (left + right) / 2

	if arr[mid] > target {
		search3(arr, target, left, mid-1)
	} else if arr[mid] < target {
		search3(arr, target, mid+1, right)
	} else {
		fmt.Println(mid)
		return mid
	}
	return -1

}

func main() {
	arr := []int{2, 4, 6, 8, 9, 11}
	target := 9
	fmt.Println(search2(arr, target))

	search3(arr, target, 0, len(arr)-1)
}
