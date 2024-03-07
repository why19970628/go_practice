package main

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	ans := make([]int, 2)
	ans[0] = searchFirst(nums, target)
	ans[1] = searchLast(nums, target)
	return ans
}

func searchFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

func searchLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				l = mid + 1
			}
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

func searchRangeV2(nums []int, target int) []int {
	ans := make([]int, 2)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	if len(nums) == 0 {
		return ans
	}
	ans[0] = searchFirstV2(nums, target)
	ans[1] = searchLastV2(nums, target)
	return ans
}

func searchFirstV2(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

func searchLastV2(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				l = mid + 1
			}
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRangeV2([]int{5, 7, 7, 8, 8, 10}, 8))
}
