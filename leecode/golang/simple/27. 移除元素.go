package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	fmt.Println(nums)
	return left
}

func removeElement111(nums []int, val int) int {
	endIdx := 0
	for _, v := range nums {
		if v != val {
			nums[endIdx] = v
			endIdx++
		}
	}
	nums = nums[:endIdx]
	return endIdx
}

func removeElement2(nums []int, val int) int {
	left := 0
	right := len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(removeElement2(nums, 2))

}
