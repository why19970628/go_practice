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

func quickSort(arr []int, first, last int) {
	flag := 3
	left := first
	right := last

	fmt.Println("flag:", flag)
	if first >= last {
		return
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		fmt.Println("flag2:", flag)
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}
	fmt.Println(arr, len(arr))

	quickSort(arr, left, flag-1)
	quickSort(arr, flag+1, right)
}

func QSort(arr []int, start int, end int) {
	var (
		key  int = arr[start]
		low  int = start
		high int = end
	)
	for {
		for low < high {
			if arr[high] < key {
				arr[low] = arr[high]
				break
			}
			high--
		}
		for low < high {
			if arr[low] > key {
				arr[high] = arr[low]
				break
			}
			low++
		}
		if low >= high {
			arr[low] = key
			break
		}
	}
	fmt.Println(arr, start)
	if low-1 > start {
		QSort(arr, start, low-1)
	}
	if high+1 < end {
		QSort(arr, high+1, end)
	}
}

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

func quickSortV3(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	pIndex := (i + j) / 2
	nums[i], nums[pIndex] = nums[pIndex], nums[i]
	p := nums[i]
	for i < j {
		for i < j && nums[j] >= p {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= p {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = p
	quickSortV3(nums, l, i-1)
	quickSortV3(nums, i+1, r)
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
func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}
	quickSortV3(testArr, 0, len(testArr)-1)

	fmt.Println(quickSortV4(testArr))
	fmt.Println(testArr)
}
