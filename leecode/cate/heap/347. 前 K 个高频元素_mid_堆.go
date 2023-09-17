package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/top-k-frequent-elements/

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


*/

func topKFrequent_innerSort(nums []int, k int) []int {
	mp := make(map[int]int, 0)
	var resp []int
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for key, _ := range mp {
		resp = append(resp, key)
	}
	fmt.Println("mp:", mp, resp)

	// 	核心思想：排序
	//	可以不用包函数，自己实现快排
	sort.Slice(resp, func(i, j int) bool {
		return mp[resp[i]] > mp[resp[j]]
	})
	fmt.Println("resp:", resp)
	return resp[:k]
}

func topKFrequent_quickSort(nums []int, k int) []int {
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	freqList := make([]int, 0, len(mp))
	for _, v := range mp {
		freqList = append(freqList, v)
	}

	quickSort(freqList, 0, len(freqList)-1)
	resp := make([]int, 0, k)
	for key, val := range mp {
		if val >= freqList[len(freqList)-k] {
			resp = append(resp, key)
		}
	}
	return resp
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pIndex := partition(nums, l, r)
	quickSort(nums, pIndex+1, r)
	quickSort(nums, l, pIndex-1)
}

func partition(nums []int, l, r int) int {
	p := nums[r]
	i, j := l, l
	for j < r {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
func main() {
	fmt.Println(topKFrequent_quickSort([]int{1, 2}, 2))
}
