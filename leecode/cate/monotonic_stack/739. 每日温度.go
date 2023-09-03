package main

import "fmt"

func dailyTemperatures(arr []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				ans = append(ans, j-i)
				break
			}
			if j == len(arr)-1 {
				ans = append(ans, 0)
			}
		}
	}
	ans = append(ans, 0)
	return ans
}

// 通常是一维数组，要寻找任一个元素的右边或者左边第一个比自己大或者小的元素的位置，此时我们就要想到可以用单调栈了。时间复杂度为O(n)。

// 暴力 O(N^2)
// 单调栈的本质是空间换时间，O(n)
// 因为在遍历的过程中需要用一个栈来记录右边第一个比当前元素高的元素，优点是整个数组只需要遍历一次。

func dailyTemperatures2(arr []int) []int {
	// 单调栈，单调递减
	stack := []int{}
	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		value := arr[i]

		// 当前元素比栈最右侧数据大，则
		for len(stack) > 0 && value > arr[stack[len(stack)-1]] {
			rightIndex := stack[len(stack)-1]
			ans[rightIndex] = i - rightIndex

			// 比当前数据小的栈顶 全部出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures3(arr []int) []int {
	ans := make([]int, len(arr))
	stack := []int{}

	for i := 0; i < len(arr); i++ {

		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {

			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}

// 单调栈 单调递增
func dailyTemperatures4(arr []int) []int {
	// 单调栈，单调递减
	var stack []int
	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		for len(stack) > 0 && val > arr[stack[len(stack)-1]] {

			// 计算差值
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 去除73
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

/*
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

#
*/
func dailyTemperaturesV5(arr []int) []int {
	resp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		resp[i] = 0
	}
	var stack []int
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {
			resp[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return resp
}
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperaturesV5(temperatures))
}
