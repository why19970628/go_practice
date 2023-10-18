package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
快速排序：分治法+递归实现
随意取一个值A，将比A大的放在A的右边，比A小的放在A的左边；然后在左边的值AA中再取一个值B，将AA中比B小的值放在B的左边，将比B大的值放在B的右边。以此类推
*/

func quick_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	rand.Seed(time.Now().Unix())
	p := rand.Intn(r-l+1) + l
	nums[r], nums[p] = nums[p], nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < nums[r] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	quick_sort(nums, l, i-1)
	quick_sort(nums, i+1, r)
}

func quickSortV2(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	pivotIndex := (i + j) / 2
	nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
	pivot := nums[i]
	for i < j {
		for i < j && pivot <= nums[j] {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quickSortV2(nums, l, i-1)
	quickSortV2(nums, i+1, r)
}

func quickSortV4(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, v := range nums[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(quickSortV4(left), append([]int{pivot}, quickSortV4(right)...)...)
}

// 快慢指针
func quickSortV3(nums []int, l, r int) {
	if l >= r {
		return
	}
	pIndex := partition(nums, l, r)
	quickSortV3(nums, l, pIndex-1)
	quickSortV3(nums, pIndex+1, r)
}

func partition(nums []int, l, r int) int {
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
	fmt.Println("swap", i, j)
	fmt.Println(nums, p)
	nums[i], nums[r] = nums[r], nums[i] // 排序
	return i                            // pivot index
}

// 快慢指针
func quickSortV5(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := partitionV5(nums, l, r)
	quickSortV5(nums, l, mid-1)
	quickSortV5(nums, mid+1, r)
}

func partitionV5(nums []int, l, r int) int {
	p := nums[r]
	i, j := l, l
	for j < r {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 0, 4}
	quickSortV5(testArr, 0, len(testArr)-1)
	fmt.Println(testArr)
}
