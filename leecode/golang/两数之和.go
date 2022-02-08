package main

import "fmt"

func twoSumSolution(nums []int, target int) {
	for i := 0; i < len(nums); i++ {
		a := i + 1
		for b := a; b < len(nums); b++ {
			if nums[i]+nums[b] == target {
				fmt.Println(i, b)
			}
		}
	}
}

func twoSumSolution2(nums []int, target int) {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	for a, b := range nums {
		if v, ok := m[target-b]; ok {
			fmt.Println(a, v)
		}
	}
}

func twoSumSolution3(nums []int, target int) {
	m := make(map[int]int)
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			fmt.Println(k, i)
		} else {
			m[v] = k
		}
	}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 13
	twoSumSolution3(nums, target)
}
