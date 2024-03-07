package main

/*
// 278. 第一个错误的版本
https://leetcode.cn/problems/first-bad-version/

*/

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func isBadVersion(num int) bool {
	return true
}
