package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//func all(nums []int)  {
//	n := [][]int{}
//	for i := 1; i < len(nums); i++ {
//
//	}
//}
func getMaxSum(arr []int) int {
	var sum, maxSum int
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//fmt.Println(maxSubArray(nums))
	fmt.Println(getMaxSum(nums))
}
