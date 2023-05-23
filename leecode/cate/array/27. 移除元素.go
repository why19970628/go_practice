package main

import "fmt"

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	endIndex := 0
	for _, v := range nums {
		if v != val {
			nums[endIndex] = v
			endIndex++
		}
	}
	nums = nums[:endIndex]
	return endIndex
}

//相向双指针法
func removeElementV2(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉

		// 找左边等于val的元素
		for left <= right && nums[left] != val {
			left++
			fmt.Println("left:", nums[left])
		}
		// 找右边不等于val的元素
		for left <= right && nums[right] == val {
			right--
			fmt.Println("right:", nums[right])
		}
		fmt.Println("-----")
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	fmt.Println(nums)
	nums = nums[:left]
	return left
}

func removeElementV22(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 等于val
		for left <= right && nums[left] != val {
			left++
		}
		// 不等于val
		for left <= right && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElementV22(nums, val))
}

//func removeElement(nums []int, val int) int {
//	endIdx := 0
//	for _, v := range nums {
//		if v != val {
//			nums[endIdx] = v
//			endIdx++
//		}
//	}
//	nums = nums[:endIdx]
//	return endIdx
//}
