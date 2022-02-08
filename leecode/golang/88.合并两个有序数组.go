package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println(nums1[m:], m)
	copy(nums1[m:], nums2)
	//nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	// 切片下标
	p1, p2 := 0, 0

	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			fmt.Println("p1: ", nums1[p1])
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			fmt.Println("p2: ", nums2[p2])
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	fmt.Println("sorted", sorted)
	copy(nums1, sorted)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge1(nums1, m, nums2, n)
	fmt.Println(nums1)
}
