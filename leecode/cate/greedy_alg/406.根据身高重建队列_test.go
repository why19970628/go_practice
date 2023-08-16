package greedy_alg

import (
	"fmt"
	"testing"
)

/*
假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。



*/

func ReconstructQueue(bills []int) bool {
	ten, five := 0, 0
	for i := 0; i < len(bills); i++ {
		pay := bills[i]
		if five < 0 {
			return false
		}
		if pay == 5 {
			five++
		} else if pay == 10 {
			ten++
			five--
		} else if pay == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func TestReconstructQueue(t *testing.T) {
	fmt.Println(ReconstructQueue([]int{5, 5, 5, 10, 20}))
}
