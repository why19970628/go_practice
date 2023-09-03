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

func TestSearchRange(t *testing.T) {
	fmt.Println(SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
