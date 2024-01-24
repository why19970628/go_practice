package main

//
//import (
//	"fmt"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//
//var target int64
//var nums []int
//
//func getMaxDigit(nums []int, target int) int {
//	for i := len(nums) - 1; i >= 0; i-- {
//		if nums[i] < target {
//			return nums[i]
//		}
//	}
//	return 0
//}
//
//func getMaxNum(nums []int, target int64) int64 {
//	var digits []int
//	numSet := make(map[int]bool)
//	for _, num := range nums {
//		numSet[num] = true
//	}
//	for target > 0 {
//		digits = append(digits, int(target%10))
//		target = target / 10
//	}
//	res := make([]int, len(digits))
//	for i := len(digits) - 1; i >= 0; i-- {
//		// 存在相同的数字，除了最后一位
//		if i > 0 && numSet[digits[i]] {
//			res[i] = digits[i]
//			continue
//		}
//		// 存在小于当前位的最大数字
//		maxDigit := getMaxDigit(nums, digits[i])
//		if maxDigit > 0 {
//			res[i] = maxDigit
//			break
//		}
//		// 需要回溯
//		for j := i + 1; j < len(digits); j++ {
//			res[j] = 0
//			maxDigit := getMaxDigit(nums, digits[j])
//			if maxDigit > 0 {
//				res[j] = maxDigit
//				break
//			}
//			if j == len(digits)-1 {
//				res = res[:len(digits)-1]
//			}
//		}
//		break
//	}
//	// 拼装目标数。
//	var sum int64 = 0
//	for i := len(res) - 1; i >= 0; i-- {
//		sum *= 10
//		if res[i] > 0 {
//			sum += int64(res[i])
//			continue
//		}
//		// 取数字数组中最大的数字。
//		sum += int64(nums[len(nums)-1])
//	}
//	return sum
//}
//
//func main() {
//	fmt.Scan(&target)
//	var input string
//	fmt.Scanln(&input)
//	numsStr := strings.Split(input, ",")
//
//	for _, numStr := range numsStr {
//		num, _ := strconv.Atoi(numStr)
//		nums = append(nums, num)
//	}
//
//	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
//	result := getMaxNum(nums, target)
//	fmt.Println(result)
//}
//
