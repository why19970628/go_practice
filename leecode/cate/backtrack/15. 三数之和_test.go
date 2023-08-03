package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumV3(nums))
}

func threeSum(nums []int) [][]int {
	resp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if i != j && i != k && j != k && nums[i]+nums[j]+nums[k] == 0 && nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					resp = append(resp, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return resp
}

// 需要去重
func threeSumV2(nums []int) [][]int {
	sort.Ints(nums)
	resp := make([][]int, 0)
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if nums[0] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			k_val := 0 - nums[j] - nums[i]
			if k, ok := mp[k_val]; ok && (i != j && i != k && j != k) && (nums[i] != nums[j] && nums[i] != k_val && nums[j] != k_val) {

				resp = append(resp, []int{nums[i], nums[j], k_val})
			}
		}
	}
	return resp
}

// 双指针 用于排完序的
func threeSumV3(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if sum == 0 {
				res = append(res, []int{n1, n2, n3})
				// l r指针移动，因为只移动一个会导致重复数据
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			}

			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}

		}

	}
	return res

}
