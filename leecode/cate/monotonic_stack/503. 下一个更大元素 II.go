package main

import "fmt"

// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
//
//数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。

//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。

func nextGreaterElementV2(nums []int) []int {
	stack := []int{}

	length := len(nums)

	ans := make([]int, length)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	for i := 0; i < length*2; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	fmt.Println(ans)
	//var stack2 []int

	//for index, v := range ans {
	//	if v == -1 {
	//
	//		stack2 = append(stack2, )
	//	}
	//}

	//var stack2 []int
	//for i := 0; i < len(nums); i++ {
	//	if len(stack2) > 0 && nums[i] < nums[stack2[len(stack2)-1]] {
	//		if ans[stack2[len(stack2)]-1] != -1 {
	//			ans[stack2[len(stack2)]-1] = nums[stack2[len(stack2)-1]]
	//		}
	//		stack2 = stack2[:len(stack2)-1]
	//	}
	//	stack2 = append(stack2, i)
	//}
	//fmt.Println(ans)
	var stack2 []int
	for i := len(nums) - 1; i >= 0; i-- {
		left := len(nums) - 1 - i
		for len(stack2) > 0 && nums[left] < nums[stack2[len(stack2)-1]] {
			if ans[i] == -1 {
				ans[stack2[len(stack2)-1]] = nums[left]
			}
			stack2 = stack2[:len(stack2)-1]
		}

		stack2 = append(stack2, left)
	}
	return ans
}

// 通过
func nextGreaterElementV22(nums []int) []int {
	stack := []int{}
	nums = append(nums, nums...)
	length := len(nums)
	ans := make([]int, length)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	for i := 0; i < length; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	ans = ans[:len(ans)/2]
	return ans
}

func main() {
	nums2 := []int{5, 4, 3, 2, 1}
	// [-1,5,5,5,5]
	fmt.Println(nextGreaterElementV22(nums2))

}
