package main

import "fmt"

// 暴力法
func trapV1(height []int) int {
	var sum int
	for index, v := range height {

		// 第一个 后一个没有面积
		if index == 0 || index == len(height)-1 {
			continue
		}

		left := height[index]  // 记录左边柱子的最高高度
		right := height[index] // 记录右边柱子的最高高度

		// 左侧
		for i := 0; i < index; i++ {
			if height[i] > left {
				left = height[i]
			}
		}
		for j := len(height) - 1; j > index; j-- {
			if height[j] > right {
				right = height[j]
			}
		}
		fmt.Println(v, left, right)

		h := min(left, right) - v
		// 找右侧比left大的
		if h > 0 {
			fmt.Println("111:", h)
			sum = sum + h
		}

	}
	return sum
}

// 双指针
func trapV2(height []int) int {
	fmt.Println(height)
	if len(height) <= 2 {
		return 0
	}
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	maxLeft[0] = height[0]
	maxRight[len(height)-1] = height[len(height)-1]

	// 记录每个柱子左边柱子最大高度
	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	// 当前列雨水面积：min(左边柱子的最高高度，记录右边柱子的最高高度) - 当前柱子高度。
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		h := min(maxLeft[i], maxRight[i]) - height[i]
		if h > 0 {
			sum += h
		}
	}
	fmt.Println(maxLeft)
	fmt.Println(maxRight)
	return sum

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(trapV2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
