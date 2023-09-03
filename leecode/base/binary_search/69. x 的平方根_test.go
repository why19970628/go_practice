package binary_search

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := -1
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}
