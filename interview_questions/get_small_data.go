package main

import "fmt"

var totalSum int

func getSmallData(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		newArr := []int{}
		for j := 0; j < i; j++ {
			if arr[j] < value {
				newArr = append(newArr, arr[j])
				totalSum += arr[j]
			}
		}
		fmt.Println(newArr)
	}
	fmt.Println(totalSum)
}

func getSmallData2(arr []int) {
	for i := 0; i < len(arr); i++ {

	}
}

func main() {
	arr := []int{1, 3, 2, 5, 4}
	getSmallData2(arr)
	fmt.Println(len(arr))
}
