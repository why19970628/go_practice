package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSortV6(nums, 0, len(nums)-1, len(nums)-k)
}

// 快慢指针
func quickSortV6(nums []int, l, r, target int) int {
	pIndex := partition2(nums, l, r)
	if pIndex == target {
		return nums[pIndex]
	}
	if pIndex > target {
		return quickSortV6(nums, l, pIndex-1, target)
	} else {
		return quickSortV6(nums, pIndex+1, r, target)
	}
}

func partition2(nums []int, l, r int) int {
	// 选择右边当做中间
	p := nums[r]
	i, j := l, l
	for j < r {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i] // 排序
	return i                            // pivot index
}

// 小根堆
func findKthLargestV2(nums []int, k int) int {
	h := &MinHeap{
		heap: make([]int, 0),
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if h.Len() < k {
			h.Insert(cur)
		} else {
			peek := h.heap[0]
			if peek < cur {
				h.Pop()
				h.Insert(cur)
			}
		}
	}
	return h.heap[0]
}

type MinHeap struct {
	heap []int
}

func (h *MinHeap) Len() int {
	return len(h.heap)
}
func (h *MinHeap) BuildHeap(array []int) {
	// copy
	h.heap = append(h.heap, array...)
	// sift down heapify
	lastNonLeafIndex := (len(array) - 2) / 2
	for i := lastNonLeafIndex; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *MinHeap) siftDown(index int) {
	leftIndex := index*2 + 1
	rightIndex := index*2 + 2
	swapIndex := index
	// 比较
	if leftIndex < len(h.heap) && h.heap[leftIndex] < h.heap[swapIndex] {
		swapIndex = leftIndex
	}
	if rightIndex < len(h.heap) && h.heap[rightIndex] < h.heap[swapIndex] {
		swapIndex = rightIndex
	}
	// 交换
	if swapIndex != index {
		h.swap(swapIndex, index)
		h.siftDown(swapIndex)
	}
}
func (h *MinHeap) siftUp(index int) {
	parentIndex := (index - 1) / 2
	swapIndex := index
	if parentIndex >= 0 && h.heap[parentIndex] > h.heap[swapIndex] {
		swapIndex = parentIndex
	}
	if swapIndex != index {
		h.swap(swapIndex, index)
		h.siftUp(swapIndex)
	}
}

func (h *MinHeap) Pop() int {
	n := len(h.heap)
	value := h.heap[0]
	h.swap(0, n-1)
	h.heap = h.heap[:n-1]
	h.siftDown(0)
	return value
}

func (h *MinHeap) Insert(num int) {
	h.heap = append(h.heap, num)
	h.siftUp(len(h.heap) - 1)
}

func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func main() {
	testArr := []int{2, 5, 3, 7, 2, 5, 8, 1, 0, 4}
	kv := findKthLargest(testArr, 2)
	fmt.Println(kv)
}
