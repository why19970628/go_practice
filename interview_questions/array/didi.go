package main

import "fmt"

// 1-100
// 这段代码首先创建一个包含1到100的切片，然后循环删除每隔2个数的元素，直到切片中只剩下一个元素。最后，它输出最后剩下的数字。

func solution2(num int) int {
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
	return index
}
func s3(n int) int {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	index := 0
	for len(nums) > 1 {
		index = (index + 2) % len(nums)
		nums = append(nums[:index], nums[index+1:]...)
	}
	fmt.Println(nums)
	return nums[0]
}

func solution1(num int) int {
	arr := make([]int, 0, num)
	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	offset := 0
	for len(arr) > 1 {
		offset = (offset + 2) % len(arr)
		arr = append(arr[:offset], arr[offset+1:]...)
	}
	fmt.Println(arr)
	return arr[0]
}
func main() {
	fmt.Println("resp:", solution1(100))
	//s3()
}
