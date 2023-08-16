package greedy_alg

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/candy/description/

func Candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	need := make([]int, len(ratings))
	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}
	sum := 0
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}

func TestCandy(t *testing.T) {
	fmt.Println(Candy([]int{1, 0, 2}))
}
