package main

/*

给你一个整数数组 nums ，请计算数组的 中心下标 。

数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。



示例 1：

输入：nums = [1,7,3,6,5,6]
输出：3
解释：
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
示例 2：

输入：nums = [1, 2, 3]
输出：-1
解释：
数组中不存在满足此条件的中心下标。
示例 3：

输入：nums = [2, 1, -1]
输出：0
解释：
中心下标是 0 。
左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。


方法一：前缀和
思路

记数组的全部元素之和为 total\textit{total}total，当遍历到第 iii 个元素时，设其左侧元素之和为 sum\textit{sum}sum，则其右侧元素之和为 total−numsi−sum\textit{total}-\textit{nums}_i-\textit{sum}total−nums
i
​
 −sum。左右侧元素相等即为 sum=total−numsi−sum\textit{sum}=\textit{total}-\textit{nums}_i-\textit{sum}sum=total−nums
i
​
 −sum，即 2×sum+numsi=total2\times\textit{sum}+\textit{nums}_i=\textit{total}2×sum+nums
i
​
 =total。

当中心下标左侧或右侧没有元素时，即为零个项相加，这在数学上称作「空和」（empty sum\text{empty sum}empty sum）。在程序设计中我们约定「空和是零」。

*/

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
