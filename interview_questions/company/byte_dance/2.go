package main

import (
	"fmt"
	"strconv"
	"strings"
)

var target int64
var nums []int
var res int64

func backtracking(target, sum int64) {
	if sum >= target {
		return
	}
	res = max(res, sum)
	for i := 0; i < len(nums); i++ {
		sum = sum*10 + int64(nums[i])
		backtracking(target, sum)
		sum = sum / 10
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&target)
	var input string
	fmt.Scanln(&input)
	numsStr := strings.Split(input, ",")

	for _, numStr := range numsStr {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	backtracking(target, 0)
	fmt.Println(res)
}
