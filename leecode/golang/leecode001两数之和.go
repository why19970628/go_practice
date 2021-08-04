package main

import "fmt"

func twoSum() {
	//dict := make(map[int]int, len(nums))
	//for i, v := range nums {
	//	fmt.Println(v)
	//	if target - v ==
	//}
}

func twoSum2(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for index, num := range nums {
		if value, ok := hashmap[num]; ok {
			nums = nums[:2]
			nums[0] = value
			nums[1] = index
			return nums
		}
		hashmap[target-num] = index
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	vkMap := make(map[int]int)
	ret := make([]int, 2)
	for k, v := range nums {
		addend := target - v
		if _, ok := vkMap[addend]; ok {
			ret[0] = vkMap[addend]
			ret[1] = k
		}
		vkMap[v] = k
	}
	fmt.Println(vkMap)
	return ret
}

func TwoSum4(arr []int, target int) []int {
	vkMap := make(map[int]int)
	ret := make([]int, 2)
	for i, v := range arr {
		vkMap[v] = i
		sub := target - v
		if _, ok := vkMap[sub]; ok {
			ret[0] = vkMap[sub]
			ret[1] = i
		}
	}
	return ret
}

func solution1(array []int, target int) {
	for i := 0; i < len(array)-1; i++ {
		b := i + 1
		for j := b; j < len(array)-1; j++ {
			if array[i]+array[j] == target {
				fmt.Println(i, j)
			}
		}
	}
}

func solution2(array []int, target int) {
	map1 := make(map[int]int, len(array))
	for i, v := range array {
		if index2, ok := map1[target-v]; ok {
			fmt.Println(index2, i)
		} else {
			map1[v] = i
		}
	}
}

func main() {
	m := []int{1, 2, 3, 4}
	res := TwoSum4(m, 7)
	fmt.Println(res)
	//nums := []int{2, 7, 9, 11, 15}
	//target := 9
}
