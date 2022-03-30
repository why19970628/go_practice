package main

import "fmt"

func main() {
	one()

}

func one() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	copy(nums1[m:], nums2[:n])
	fmt.Println(nums1)

}
