package binary_search

import (
	"fmt"
	"testing"
)

func SearchRange(nums []int, target int) []int {
	resp := make([]int, 2)
	resp[0] = -1
	resp[1] = -1
	left := 0
	right := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if nums[mid+1] == target {
				resp[0] = mid
				resp[1] = mid + 1
			}
			if nums[mid-1] == target {
				resp[0] = mid - 1
				resp[1] = mid
			}
			return resp
		}
		if nums[mid] < target {
			right = mid + 1
		}
		if nums[mid] > target {
			left = mid - 1
		}
	}
	return resp
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 查找
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 查找
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

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

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{2, 2}, 2))
}
