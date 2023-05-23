package main

import (
	"fmt"
	"math"
)

//
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum = sum + nums[j]
			if sum >= target {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	return ans
}

func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			data := nums[j]
			sum = sum + data
			if sum >= target {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if math.MaxInt32 == ans {
		return 0
	}
	return ans
}

// https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/305704/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio/
// 滑动窗口
func minSubArrayLen3(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minSubArrayLen2(11, []int{1, 2, 3, 4, 5}))
}
