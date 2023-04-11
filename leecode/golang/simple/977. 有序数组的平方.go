package main

import "fmt"

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	length := len(nums) - 1
	ans := make([]int, len(nums))
	for l <= r {
		ldata := nums[l] * nums[l]
		rdata := nums[r] * nums[r]
		if ldata > rdata {
			ans[length] = ldata
			l++
		} else {
			ans[length] = rdata
			r--
		}

		length--
	}
	return ans

}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
