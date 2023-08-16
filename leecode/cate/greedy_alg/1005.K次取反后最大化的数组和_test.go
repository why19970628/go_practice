package greedy_alg

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func LargestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			k--
			nums[i] = -nums[i]
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func LargestSumAfterKNegationsV2(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < nums[i]; i++ {
		if k > 0 && nums[i] < 0 {
			k--
			nums[i] = -nums[i]
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result

}

func TestLargestSumAfterKNegations(t *testing.T) {
	fmt.Println(LargestSumAfterKNegations([]int{4, 2, 3}, 1))
}
