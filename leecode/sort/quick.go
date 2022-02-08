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

// 选定一个基准值(任意选,以第一个为例)
//定义左右指针指向左右两端
//左指针往右移动,如果遇到大于基准值的数就把它和右指针的值调换位置,然后左指针不动,右指针开始向左移动,如果遇到小于基准值的数就把他和左指针的值调换位置,然后开始移动左指针,以此类推,知道左右指针相遇
//递归上述过程直到排序结束
//func QuickSort(intList []int) {
//	// 如果长度小于等于1就直接结束
//	if len(intList) <= 1 {
//		return
//	}
//	// 1. 将第一个值选定为基准值
//	flag := intList[3]
//	// 定义左右指针
//	left, right := 0, len(intList)-1
//
//	for i := 1; i <= right; {
//		fmt.Println(i)
//		if intList[i] > flag {
//			intList[i], intList[right] = intList[right], intList[i]
//			right--
//		} else {
//			intList[i], intList[left] = intList[left], intList[i]
//			i++
//			left++
//		}
//	}
//	fmt.Println("intList0:", intList, left, right, "flag:", flag)
//	//fmt.Println("intList[:left]", intList[:left])
//	// 递归
//	QuickSort(intList[:left])
//
//	fmt.Println("intList1:", intList)
//	QuickSort(intList[left+1:])
//	fmt.Println("intList2:", intList)
//}

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

func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}
	//quickSort(testArr)
	quickSort(testArr, 0, len(testArr)-1)
	fmt.Println(testArr)
	//quick_sort(testArr, 0, len(testArr)-1)
	//fmt.Println(testArr)
}
