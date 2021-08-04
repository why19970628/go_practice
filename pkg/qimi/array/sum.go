package main

import "fmt"

func arraySum(array []int) {
	sum := 0
	for _, value := range array {
		sum = sum + value
	}
	fmt.Println(sum)
}

func arraySumIndex(array []int, sum int) {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i]+array[j] == sum {
				fmt.Println(i, j)
			}
		}
	}
}

func main() {
	var array = []int{10, 20, 30}
	arraySum(array)

	var array2 = []int{1, 3, 5, 7, 8}
	sum := 8
	arraySumIndex(array2, sum)
}
