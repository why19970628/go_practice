package greedy_alg

import (
	"fmt"
	"sort"
	"testing"
)

//排序后，局部最优
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var result = 0
	gLength := len(g) - 1
	for i := len(s) - 1; i >= 0; i-- {
		for ; gLength >= 0; gLength-- {
			if g[gLength] <= s[i] {
				result++
				gLength--
				break
			}
		}

	}

	return result
	// 从小到大
	//child := 0
	//for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
	//	if s[sIdx] >= g[child] {//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
	//		child++
	//	}
	//}

	//return child
}

func TestPrintV4(t *testing.T) {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}
