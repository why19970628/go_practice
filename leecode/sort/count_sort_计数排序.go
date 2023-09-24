package main

import (
	"fmt"
	"math"
)

// 计数排序

// 都是使用桶 bucket (额外的空间)将有限范围(比如100) arr 映射到一个数组里面

// 然后遍历 bucket 拿到数据次数依次摆放

func CountingSort(arr []int) {

	max := math.MinInt64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	buckets := make([]int, max+1)
	for _, v := range arr {
		buckets[v] += 1
	}

	index := 0
	for element, count := range buckets {
		for count > 0 && index < len(arr) {
			arr[index] = element
			index++
			count--
		}
	}

}

func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 0, 4}
	CountingSort(testArr)
	fmt.Println(testArr)
}
