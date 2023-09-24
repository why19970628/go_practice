package main

import "fmt"

func BucketSort(arr []int) {

	//1.找到最大值和最小值
	max := arr[0]
	min := arr[0]
	mp := make(map[int]int)
	for _, element := range arr {
		if element > max {
			max = element
		}
		if element < min {
			min = element
		}
		mp[element]++
	}
	buckets := make([][]int, max+1)
	for element, count := range mp {
		for count > 0 {
			buckets[element] = append(buckets[element], element)
			count--
		}
	}

	index := 0
	for _, bucket := range buckets {
		for i := 0; i < len(bucket); i++ {
			arr[index] = bucket[i]
			index++
		}
	}
}

func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 0, 4}
	BucketSort(testArr)
	fmt.Println(testArr)
}
