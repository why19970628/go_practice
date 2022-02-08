package main

import "fmt"

func permute(nums []int) [][]int {
	res = [][]int{}
	run(nums, []int{})
	return res
}

var res [][]int

func run(nums []int, cur []int) {
	if len(cur) == len(nums) {
		fmt.Println("append::::::::", cur)
		res = append(res, cur)
		return
	}

	fmt.Println("cur2:", cur)

	for i := 0; i < len(nums); i++ {
		exist := false
		for j := 0; j < len(cur); j++ {
			if nums[i] == cur[j] {
				exist = true
				break
			}
		}
		if exist {
			continue
		}
		fmt.Println("bbb", nums, cur, append(cur, nums[i]))
		run(nums, append(cur, nums[i]))
	}
}

func main() {
	num := []int{1, 2, 3}
	permute(num)
	fmt.Println(res)
}
