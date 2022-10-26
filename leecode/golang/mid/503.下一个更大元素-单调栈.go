package main

import "fmt"

func main() {
	arr := []int{2, 6, 3, 8, 10, 9}
	res := nextGreaterElement2(arr)
	fmt.Println(res)
}

func nextGreaterElement2(arr []int) []int {
	ans := make([]int, len(arr))
	// 单调栈，单调递减
	tmp := make([]int, 0)
	for index, value := range arr {
		for len(tmp) > 0 && value > arr[tmp[len(tmp)-1]] {
			// 当前栈钉元素 index
			rightIndex := tmp[len(tmp)-1]
			fmt.Println(arr[rightIndex], rightIndex)
			ans[rightIndex] = value
			tmp = tmp[:len(tmp)-1]
		}
		//
		tmp = append(tmp, index)
	}
	return ans
}

func nextGreaterElement(arr []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		a := false
		for j := i + 1; j < len(arr); j++ {
			if j == len(arr)-1 {
				continue
			}
			if arr[i] < arr[j] {
				fmt.Println(arr[j])
				res = append(res, arr[j])
				a = true
				break
			}
		}
		if a == false {
			res = append(res, -1)
		}
	}
	fmt.Println(res)
	return res
}
