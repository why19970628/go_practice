package main

import (
	"container/heap"
)

//方法一：小顶堆
func topKFrequentV1(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

//构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 实际实现快排
func findTopKFrequentV3(nums []int, k int) []int {
	// 统计元素的频率
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// 将频率信息存储到切片中
	freqList := make([]int, 0, len(frequency))
	for _, freq := range frequency {
		freqList = append(freqList, freq)
	}

	// 利用快速选择算法找到第 k 高的频率
	//kthFrequency := quickSelect(freqList, k)
	//
	//// 构建结果列表
	result := make([]int, 0, k)
	//for num, freq := range frequency {
	//	if freq == kthFrequency {
	//		result = append(result, num)
	//	}
	//}

	return result
}

//func quickSelect(arr []int, k int) int {
//	left, right := 0, len(arr)-1
//	for left <= right {
//		pivotIndex := partition(arr, left, right)
//		if pivotIndex == k-1 {
//			return arr[pivotIndex]
//		} else if pivotIndex < k-1 {
//			left = pivotIndex + 1
//		} else {
//			right = pivotIndex - 1
//		}
//	}
//	return -1 // 如果 k 超出范围，返回一个无效值
//}
//
//func partition(arr []int, left, right int) int {
//	pivot := arr[right]
//	i := left
//	for j := left; j < right; j++ {
//		if arr[j] >= pivot {
//			arr[i], arr[j] = arr[j], arr[i]
//			i++
//		}
//	}
//	arr[i], arr[right] = arr[right], arr[i]
//	return i
//}

func main() {
	//nums := []int{1, 1, 1, 2, 2, 3}
	//k := 2
	//result := topKFrequentV2([]int{3, 0, 1, 0}, 1)
	//fmt.Println(result) // 输出 [1, 2]
}
