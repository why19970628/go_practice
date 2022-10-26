package main

import "fmt"

//dailyTemperatures 单调栈
func dailyTemperatures(arr []int) []int {
	ans := make([]int, len(arr))
	// 单调栈，单调递减
	tmp := make([]int, 0)
	for index, value := range arr {
		// 如果当前元素比栈元素大，
		// 则 记录距离 栈弹出元素
		for len(tmp) > 0 && value > arr[tmp[len(tmp)-1]] {
			rightIndex := tmp[len(tmp)-1]
			ans[rightIndex] = index - rightIndex
			tmp = tmp[:len(tmp)-1]
		}
		// 如果元素比栈顶小 单调栈加
		tmp = append(tmp, index)

	}
	return ans
}

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(arr))
}
