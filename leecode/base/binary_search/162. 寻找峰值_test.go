package main

func findPeakElement(nums []int) int {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}
