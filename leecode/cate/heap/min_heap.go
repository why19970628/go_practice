package main

//
//type MinHeap struct {
//	heap []int
//}
///*
//
//https://tianjietalk.github.io/#/leetcode/heap
//*/
//
//func (h *MinHeap) BuildHeap(array []int) {
//	// copy
//	h.heap = append(h.heap, array...)
//	// sift down heapify
//	lastNonLeafIndex := (len(array) - 2) / 2
//	for i := lastNonLeafIndex; i >= 0; i-- {
//		h.siftDown(i)
//	}
//}
//
//func (h *MinHeap) siftDown(index int) {
//	leftIndex := index*2 + 1
//	rightIndex := index*2 + 2
//	swapIndex := index
//	// 比较
//	if leftIndex < len(h.heap) && h.heap[leftIndex] < h.heap[swapIndex] {
//		swapIndex = leftIndex
//	}
//	if rightIndex < len(h.heap) && h.heap[rightIndex] < h.heap[swapIndex] {
//		swapIndex = rightIndex
//	}
//	// 交换
//	if swapIndex != index {
//		h.swap(swapIndex, index)
//		h.siftDown(swapIndex)
//	}
//}
//func (h *MinHeap) siftUp(index int) {
//	parentIndex := (index - 1) / 2
//	swapIndex := index
//	if parentIndex >= 0 && h.heap[parentIndex] > h.heap[swapIndex] {
//		swapIndex = parentIndex
//	}
//	if swapIndex != index {
//		h.swap(swapIndex, index)
//		h.siftUp(swapIndex)
//	}
//}
//
//func (h *MinHeap) Pop() int {
//	n := len(h.heap)
//	value := h.heap[0]
//	h.swap(0, n-1)
//	h.heap = h.heap[:n-1]
//	h.siftDown(0)
//	return value
//}
//
//func (h *MinHeap) Insert(num int) {
//	h.heap = append(h.heap, num)
//	h.siftUp(len(h.heap) - 1)
//}
//
//func (h *MinHeap) swap(i, j int) {
//	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
//}
