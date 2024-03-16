package main

import (
	"container/heap"
	"fmt"
)

/*
https://leetcode.cn/problems/kth-largest-element-in-an-array/

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/
func findKthLargest(nums []int, k int) int {
	return quickSortV2(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSortV2(nums []int, l, r, target int) int {
	pIndex := partition2(nums, l, r)
	fmt.Println("pIndex:", pIndex)
	if target == pIndex {
		return nums[target]
	}
	if target > pIndex {
		return quickSortV2(nums, pIndex+1, r, target)
	} else {
		return quickSortV2(nums, l, pIndex-1, target)
	}
}

func partition2(nums []int, l, r int) int {
	p := nums[r]
	i, j := l, l
	for j < len(nums) {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

// 堆排序 小顶堆
type iheap []int

func (h iheap) Len() int {
	return len(h)
}

func (h iheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h iheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *iheap) Push(x any) {
	*h = append(*h, x.(int))
}

// 弹出数组最后一个元素，同时数组长度--，堆内部通过替换堆顶元素来进行 heapifyDown 调整
func (h *iheap) Pop() any {
	n := h.Len()
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func (h iheap) Peek() any {
	return h[0]
}

/*
方法二： 维护一个长度为k的小顶堆，利用堆的特性每次入堆进行heapifyUp调整，堆顶元素就是第K大
*/
func findKthLargestV2(nums []int, k int) int {
	h := iheap{}
	heap.Init(&h)
	for i := 0; i < len(nums); i++ {
		if h.Len() < k {
			heap.Push(&h, nums[i])
			fmt.Println("push:", h)
		} else {
			// //小顶堆 堆顶元素小于当前元素
			top := h.Peek().(int)
			if nums[i] > top {
				_ = heap.Pop(&h)
				heap.Push(&h, nums[i])
			}
		}

		// v2版本
		//heap.Push(&h, nums[i])
		//if h.Len() > k {
		//	heap.Pop(&h)
		//}
	}
	fmt.Println(h)
	return heap.Pop(&h).(int)
}

func main() {
	fmt.Println(findKthLargestV2([]int{-1, 2, 0}, 2))
	//fmt.Println(findKthLargestV2([]int{3, 2, 1, 5, 6, 4}, 2))
}
