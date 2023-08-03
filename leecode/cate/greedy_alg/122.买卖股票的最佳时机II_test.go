package greedy_alg

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var sum int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
