package main

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/container-with-most-water/

func maxArea(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxAreaV2(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r++
		}
	}
	return ans
}

func TestSearchRange(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
