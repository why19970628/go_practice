package main

import "fmt"

// 冒泡排序
/*
1.冒泡排序
基本思想：每一次把最小的元素放到最前面，循环n-1次，为什么循环n-1次，因为最后一个元素不用比较，肯定是最大的。
*/
var array3 []int = []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}

func main() {
	BubbleSortV2()
}

func BubbleSortV2() {
	for i := 0; i < len(array3); i++ {
		for j := 0; j < len(array3)-1; j++ {
			if array3[i] < array3[j] {
				array3[i], array3[j] = array3[j], array3[i]
			}
		}
	}
	fmt.Println(array3)
}
