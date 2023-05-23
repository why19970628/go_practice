package main

import "fmt"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/

// 暴力法
func minSubArrayLen(target int, nums []int) int {
	//ans := make([]int, 0)
	minLenght := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		count := 0

		//a := 0
		for j := i; j < len(nums); j++ {
			count++
			sum = sum + nums[j]

			//fmt.Println("sum:", sum)
			if sum >= target {
				if minLenght == 0 {
					minLenght = count
				} else {
					minLenght = min(minLenght, count)
				}
				break
			}

		}
	}
	return minLenght
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 滑动窗口!
func minSubArrayLenv2(target int, nums []int) int {
	left := 0 //左节点
	n := len(nums)
	sum := 0
	res := n + 1
	// 右节点
	for right, v := range nums {
		sum += v
		for sum >= target {
			sum = sum - nums[left]
			res = min(res, right-left+1)
			left++
		}
	}
	if res == n+1 {
		return 0
	}
	return res

}

//链接：https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/1959532/biao-ti-xia-biao-zong-suan-cuo-qing-kan-k81nh/
// 滑动窗口
func minSubArrayLenV22(target int, nums []int) int {
	n := len(nums)
	ans, sum, left := n+1, 0, 0
	for right, v := range nums {
		sum += v
		for sum >= target { // 满足要求
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

func main() {
	fmt.Println(minSubArrayLenv2(15, []int{1, 2, 3, 4, 5}))
}
