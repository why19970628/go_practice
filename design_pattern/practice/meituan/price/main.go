package main

import (
	"fmt"
)

func main() {
	var cf CashFactory
	var price Pricer
	price = cf.GetCashPrice(ORIGINAL_PRICE)
	fmt.Println(price.Price(210))

	price = cf.GetCashPrice(TEN_PERCENT_OFF)
	fmt.Println(price.Price(210))

	price = cf.GetCashPrice(FULL_100_MINUS_20)
	fmt.Println(price.Price(210))
}
