package main

import "fmt"

// 1-100
// 这段代码首先创建一个包含1到100的切片，然后循环删除每隔2个数的元素，直到切片中只剩下一个元素。最后，它输出最后剩下的数字。

func solution2(num int) {
	arr := make([]int, 0, 100)
	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	index := 2
	for len(arr) > 1 {
		if index > len(arr)-1 {
			index = index % len(arr)
		}
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("arr:", index, arr)
		index += 2
	}
}
func s3() {
	nums := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}

	index := 0
	for len(nums) > 1 {
		index = (index + 2) % len(nums)
		nums = append(nums[:index], nums[index+1:]...)
	}
	fmt.Println(nums)
}
func main() {
	solution2(100)
	//s3()
}
