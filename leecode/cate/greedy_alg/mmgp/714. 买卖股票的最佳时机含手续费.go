package mmgp

func maxProfit_714(prices []int, fee int) int {
	var sum int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1]-fee > 0 {
			sum += prices[i] - prices[i-1] - fee
		}
	}
	return sum
}
