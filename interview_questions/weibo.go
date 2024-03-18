package main

/*
一个严格有序的数组，比如[1,2,3,4,5,6]，

经过一次翻转后变成部分有序，比如[3,4,5,6,1,2]，
求翻转后的数组的最小值的序号。输入是翻转后的数据，输出是最小值的序号。
*/
func solution(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
