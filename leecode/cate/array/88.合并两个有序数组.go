package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	temp := make([]int, 0, m+n)
	for {
		fmt.Println(temp)
		if i == m {
			temp = append(temp, nums2[j:]...)
			break
		}
		if j == n {
			temp = append(temp, nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}

	}
	fmt.Println("sorted", temp)
	copy(nums1, temp)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
