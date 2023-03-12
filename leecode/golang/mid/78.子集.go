package main

import "fmt"

func subset(arr []int) (res [][]int) {
	length := len(arr)
	if length == 0 {
		return nil
	}
	ans := make([][]int, 0, 1<<length)
	var f func(i int)
	path := make([]int, 0, length)
	f = func(i int) {
		ans = append(ans, append([]int(nil), path...)) // 固定答案
		if i == length {
			return
		}
		for j := i; i < length; j++ {
			path = append(path, arr[j])
			f(i + 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	f(0)
	return ans
}

func main() {
	res := subset([]int{1, 2, 3})
	fmt.Println(res)
}
