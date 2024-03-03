package main

import "fmt"

/*
https://leetcode.cn/problems/move-zeroes/

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。



示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]

*/

// 快慢指针
func moveZeroes(nums []int) {
	slow := 0
	fast := 1
	for slow < fast && fast < len(nums) {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] != 0 {
			slow++
		}
		fast++
	}

}

// 输入: nums = [0,1,0,3,12]
func moveZeroesV2(nums []int) {
	slow, fast := 0, 1
	for slow < fast && fast < len(nums) {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] != 0 {
			slow++
		}
		fast++
	}
}

// 输入: nums = [0,1,0,3,12]
func moveZeroesV3(nums []int) {
	slow, fast := 0, 1
	for slow < fast && fast < len(nums) {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] != 0 {
			slow++
		}
		fast++
	}
}

func MoveTargetToFront(nums []int, target int) {
	slow, fast := 0, 1
	for slow < fast && fast < len(nums) {
		if nums[fast] == target && nums[slow] != target {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] == target {
			slow++
		}
		fast++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroesV3(nums)
	fmt.Println(nums)
}
