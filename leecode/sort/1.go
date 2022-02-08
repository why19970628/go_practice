package main

import "fmt"

// 梳子排序
func ComoSort(arr []int) []int {

	// 数组长度
	length := len(arr)

	// 初始化步长
	gap := length

	for gap > 1 {

		// 递减率设定为1.3，这是经过实验证明的最快的
		gap = gap * 10 / 13
		fmt.Println("gap:", gap)

		// 收缩
		for i := 0; i+gap < length; i++ {
			fmt.Println("i:", i)
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 5, 5, 9, 8, 8, 6, 4, 4, 9, 3, 2, 7, 8, 5, 1, 2}
	fmt.Println("arr:", len(arr))
	fmt.Println(ComoSort(arr))
}
