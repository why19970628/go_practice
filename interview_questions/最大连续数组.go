package main

import (
	"fmt"
	"math"
)

//
func main() {
	var numArr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//var numArr = []int{-2, 1, -3, 1, -3, 4, 2, -5, 4}
	//totalNum := maxAddArray(numArr)
	//fmt.Println("当前数组中，最大之和为:", totalNum)
	//maxAddArray(numArr)

	//fmt.Println(maxSubArray(numArr))
	fmt.Println(maxSubArray2(numArr))
}

func maxSubArray2(nums []int) int {
	//var numArr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println("v:", nums[i])
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		fmt.Println("curr:", nums[i])

		if nums[i] > maxSum {
			maxSum = nums[i]
		}
		fmt.Println("max:", maxSum)
		fmt.Println()

	}

	return maxSum
}

// 动态规划

//这里使用max记录最大值
//tmp记录当前值，如果当前值已经小于0，可以直接抛弃
//如果大于0，尝试更新max
func plan2(nums []int) int {
	//var numArr = []int{-2,1,-3,4,-1,2,1,-5,4}
	//max := int(^(int(^uint(0) >> 1)))
	max := math.MinInt64
	fmt.Println(max)
	tmp := 0
	for _, v := range nums {
		fmt.Println("v :", v, "tmp:", tmp)
		if tmp > 0 {
			tmp += v
		} else {
			tmp = v
		}
		if max < tmp {
			max = tmp
		}
		fmt.Println("max :", max)
		fmt.Println()
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力法
func maxAddArray(num []int) {
	max := 0
	all := make([][]int, 0)
	for j := 0; j < len(num); j++ {
		for i := j; i < len(num); i++ {
			a := []int{}
			a = append(a, num[j:i+1]...)
			all = append(all, a)
		}
		fmt.Println(all)
		fmt.Println()
	}
	for _, v := range all {
		if getArrSum(v) > max {
			max = getArrSum(v)
		}
	}
	fmt.Println(max)
}

func getArrSum(a []int) int {
	b := 0
	for _, v := range a {
		b = b + v
	}
	return b
}

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max_sum {
			max_sum = nums[i]
		}
		fmt.Println(max_sum)
	}
	return max_sum
}
