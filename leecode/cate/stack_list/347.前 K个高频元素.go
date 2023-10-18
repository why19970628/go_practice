package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// https://leetcode.cn/problems/top-k-frequent-elements/solutions/1760658/by-carlsun-2-hybi/

//方法二:利用O(nlogn)排序
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}

	ans := make([]int, 0)
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return map_num[ans[i]] > map_num[ans[j]]
	})
	return ans[:k]
}

//方法一：小顶堆
// 时间复杂度：O(nlogk)
//func topKFrequent2(nums []int, k int) []int {
//	map_num := map[int]int{}
//	//记录每个元素出现的次数
//	for _, item := range nums {
//		map_num[item]++
//	}
//	h := &IHeap{}
//	heap.Init(h)
//	//所有元素入堆，堆的长度为k
//	for key, value := range map_num {
//		heap.Push(h, [2]int{key, value})
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//	res := make([]int, k)
//	//按顺序返回堆中的元素
//	for i := 0; i < k; i++ {
//		res[k-i-1] = heap.Pop(h).([2]int)[0]
//	}
//	return res
//}
//
////构建小顶堆
//type IHeap [][2]int
//
//func (h IHeap) Len() int {
//	return len(h)
//}
//
//func (h IHeap) Less(i, j int) bool {
//	return h[i][1] < h[j][1]
//}
//
//func (h IHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *IHeap) Push(x interface{}) {
//	*h = append(*h, x.([2]int))
//}
//
//func (h *IHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	fmt.Println("old:", old)
//	// 取最右侧 最后一个
//	x := old[n-1]
//	// 数组去除最后一个
//	*h = old[0 : n-1]
//	return x
//}

type IHeap [][2]int

func (I IHeap) Len() int {
	return len(I)
}

func (I IHeap) Less(i, j int) bool {
	return I[i][1] < I[j][1]

}

func (I IHeap) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func (I *IHeap) Push(x interface{}) {
	*I = append(*I, x.([2]int))
}

func (I *IHeap) Pop() interface{} {
	old := *I
	n := len(old)
	x := old[n-1]
	*I = old[:n-1]
	return x

}

func topKFrequentHeap(nums []int, k int) []int {
	mapNum := make(map[int]int)
	for _, item := range nums {
		mapNum[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	for num, count := range mapNum {
		heap.Push(h, [2]int{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	resp := make([]int, k)
	for i := 0; i < k; i++ {
		resp[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return resp
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequentHeap(nums, k))
}
