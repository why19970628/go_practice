package main

import "fmt"

/*
2.归并排序
思路
归并排序利用了分治的思想来对序列进行排序。对一个长为 n 的待排序的序列，我们将其分解成两个长度为 n/2的子序列。每次先递归调用函数使两个子序列有序，然后我们再线性合并两个有序的子序列使整个序列有序。

求mid，递归排左边、右边，合并
*/
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
	fmt.Println(arr)

	fmt.Println("-----")
}

func merge(arr []int, start, mid, end int) {
	var tmparr = []int{}
	var s1, s2 = start, mid + 1
	fmt.Println("s1, s2", s1, s2, " mid, end:", mid, end)
	for s1 <= mid && s2 <= end {
		if arr[s1] > arr[s2] {
			tmparr = append(tmparr, arr[s2])
			s2++
		} else {
			tmparr = append(tmparr, arr[s1])
			s1++
		}
	}
	fmt.Println("tmparr0:", tmparr, s1, s2, "mid", mid, end)
	//上述步骤肯定有一堆先拿完 故将剩下的一堆直接加入到tempArr
	if s1 <= mid {
		fmt.Println("s1 add:", s1, arr[s1:mid+1])
		tmparr = append(tmparr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		fmt.Println("s2 add:", s2, arr[s2:end+1])
		tmparr = append(tmparr, arr[s2:end+1]...)
	}
	fmt.Println("tmparr1:", tmparr)
	for pos, item := range tmparr {
		arr[start+pos] = item
	}
}

var a = []int{4, 3, 1, 0, 8, 7, 6, 5}

func main() {
	mergeSort(a, 0, len(a)-1)
	fmt.Println("a:", a)
}
