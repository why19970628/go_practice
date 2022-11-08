package main

import "fmt"

// o(n)
func searchInsert(nums []int, target int) int {
	res := 0
	for i, v := range nums {
		if v < target {
			res = i + 1
			continue
		}
		if v == target {
			res = i
			return res
		}
		if v > target {
			return res
		}
	}
	return len(nums)
}

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//请必须使用时间复杂度为 O(log n) 的算法。
//
//
// O(lognN)
func searchInsert2(nums []int, target int) int {
	res := len(nums)
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4

	fmt.Println("====================")

	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 7)) // 4

}
