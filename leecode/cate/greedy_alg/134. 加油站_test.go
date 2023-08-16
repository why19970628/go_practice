package greedy_alg

import (
	"fmt"
	"testing"
)

func CanCompleteCircuit(gas []int, cost []int) int {
	currSum := 0
	totalSum := 0
	startIndex := 0
	for i := 0; i < len(gas); i++ {
		left := gas[i] - cost[i]
		totalSum += left
		currSum += left
		// 不够，则从此坐标开始
		if currSum < 0 {
			startIndex = i + 1
			currSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return startIndex
}

func TestCanCompleteCircuit(t *testing.T) {
	fmt.Println(CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
