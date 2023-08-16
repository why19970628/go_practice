package greedy_alg

import (
	"fmt"
	"testing"
)

func LemonadeChange(bills []int) bool {
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

func TestLemonadeChange4(t *testing.T) {
	fmt.Println(LemonadeChange([]int{5, 5, 5, 10, 20}))
}
