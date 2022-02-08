package main

import "fmt"

// https://zhuanlan.zhihu.com/p/125556928
//桶排序
//数组下标为排序元素的值，数组值为元素出现的个数
func solution(arr []int) {
	m := make([]int, 10)
	//for i := 0; i < len(arr)-1; i++ {
	//	m[arr[i]]++
	//}

	for _, v := range arr {
		m[v]++
	}

	var newArr []int
	for index, val := range m {
		for val > 0 {
			newArr = append(newArr, index)
			val--

		}
	}
	fmt.Println(m)
	fmt.Println(newArr, len(newArr))
}

func main() {
	arr := []int{1, 5, 5, 9, 8, 8, 6, 4, 4, 9, 3, 2, 7, 8, 5, 1, 2}
	fmt.Println(len(arr))
	solution(arr)
}
