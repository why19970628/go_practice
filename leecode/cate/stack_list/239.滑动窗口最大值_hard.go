package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回 滑动窗口中的最大值 。

// https://leetcode.cn/problems/sliding-window-maximum/
// n*k
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	res := make([]int, 0)
	n := len(nums)
	for i := 0; i < n-k+1; i++ {
		res = append(res, max(nums[i:i+k]))
	}
	return res
}

func max(numList []int) int {
	maxNum := numList[0] //假设最大值是第一个
	for _, v := range numList {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// 优先队列  大顶堆-------------------------------------------------------
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	fmt.Println("q:", q)
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	fmt.Println("ans:", ans)
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// 优先队列  大顶堆-------------------------------------------------------

// https://leetcode.cn/problems/sliding-window-maximum/solutions/2126542/ji-lu-shu-zu-yuan-su-zhong-de-zhi-yu-xia-naaz/
func main() {
	n := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow2(n, k))
}
