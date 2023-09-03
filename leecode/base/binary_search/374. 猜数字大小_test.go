package binary_search

func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		}
		if guess(mid) == 1 {
			left = mid + 1
		}
		// -1：我选出的数字比你猜的数字小 pick < num
		if guess(mid) == -1 {
			right = mid - 1
		}
	}
	return -1
}

func guess(num int) int {
	return num
}
