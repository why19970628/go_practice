package main

import (
	"fmt"
	"testing"
	"unicode"
)

/*
目标移动
【描述】
给定一个数组 nums 以及一个整数 target 。
你需要把数组中等于target的元素移动到数组的最前面，并且其余的元素相对顺序不变。
你的所有移动操作都应该在原数组上面操作。
数组的长度范围是: [1, 100000]
如果数组中没有出现 target 则不需要对原数组进行修改。
*/
func MoveTargetToFront(nums []int, target int) {
	slow, fast := 0, 1
	for slow < fast && fast < len(nums) {
		if nums[fast] == target && nums[slow] != target {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] == target {
			slow++
		}
		fast++
	}
}

func TestMoveTargetToFront(t *testing.T) {
	nums := []int{2, 1, 4, 2, 1, 2, 3}
	MoveTargetToFront(nums, 2)
	fmt.Println(nums)
}

/*
在LeetCode上，这个问题是第 #1160 题，“拼写单词”（Find Words That Can Be Formed by Characters）的变种。原问题是找出给定单词数组中可以由字符表中的字符拼写出的单词的总长度。
给定一个字符串str，返回字符串中字母顺序最大的而且同时在字符串中出现大写和小写的字母。 如果不存在这样的字母，返回‘~‘。
请返回大写字母
 |str|<=1000
*/

func FindMaxUpper(str string) byte {
	maxUpper := '~'
	lowerMap := make(map[rune]bool)
	upperMap := make(map[rune]bool)

	for _, char := range str {
		if unicode.IsUpper(char) {
			upperMap[char] = true
			lowerChar := unicode.ToLower(char)
			if lowerMap[lowerChar] {
				if maxUpper == '~' || char > maxUpper {
					maxUpper = char
				}
			}
		} else if unicode.IsLower(char) {
			lowerMap[char] = true
			upperChar := unicode.ToUpper(char)
			if upperMap[upperChar] {
				if maxUpper == '~' || upperChar > maxUpper {
					maxUpper = upperChar
				}
			}
		}
	}

	if maxUpper == '~' {
		return '~'
	}

	return byte(maxUpper)
}

func TestFindMaxUpper(t *testing.T) {
	fmt.Println(string(FindMaxUpper("aBCA")))
}

/*
931. 下降路径最小和
https://leetcode.cn/problems/minimum-falling-path-sum/description/
给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。
下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列。

[

	[1,2,3],
	[4,5,6],
	[7,8,9]

]
*/
func minFallingPathSum(A [][]int) int {
	n := len(A)
	m := len(A[0])
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				A[i][j] += min(A[i-1][j], A[i-1][j+1])
			} else if j == m-1 {
				A[i][j] += min(A[i-1][j-1], A[i-1][j])
			} else {
				A[i][j] += min(A[i-1][j-1], min(A[i-1][j], A[i-1][j+1]))
			}
		}
	}
	minSum := A[n-1][0]
	for j := 1; j < m; j++ {
		minSum = min(minSum, A[n-1][j])
	}
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinFallingPathSum(t *testing.T) {
	fmt.Println(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
