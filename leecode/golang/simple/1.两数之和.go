package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(F1(nums, target))
	fmt.Println(F2(nums, target))
	var a int = 1000
	var b int = 1000
	fmt.Println("a==b", a == b)

	type pos [2]int
	a1 := pos{4, 5}
	a2 := pos{4, 5}
	fmt.Println(a1 == a2)

}

func F1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; i < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func F2(nums []int, target int) []int {
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]] = i
	}
	fmt.Println(temp)
	for i := 0; i < len(nums); i++ {
		if v, ok := temp[target-nums[i]]; ok {
			return []int{i, v}
		}
	}

	return nil
}
