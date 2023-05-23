package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	mp := map[int]int{}
	for i, v := range nums1 {
		mp[v] = i
	}
	// 单调栈
	stack := []int{0}

	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			// 要移除的的栈顶index
			top := stack[len(stack)-1]
			if _, ok := mp[nums2[top]]; ok { // 看map里是否存在这个元素
				index := mp[nums2[top]] // 根据map找到nums2[top] 在 nums1中的下标
				res[index] = nums2[i]
			}

			// 比当前数据小的栈顶 全部出栈
			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, i)
	}
	return res
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	// 单调栈
	stack := []int{}

	mp := make(map[int]int)
	for i, v := range nums1 {
		mp[v] = i
	}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			rightIndex := stack[len(stack)-1]
			rightValue := nums2[rightIndex]
			if n1Index, ok := mp[rightValue]; ok {
				res[n1Index] = nums2[i]
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	return res
}

func nextGreaterElement3(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mp[nums1[i]] = i
	}
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			if index, ok := mp[nums2[stack[len(stack)-1]]]; ok {
				ans[index] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return ans
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement3(nums1, nums2))
}
