package main

const (
	ORIGINAL_PRICE    = iota // 0
	TEN_PERCENT_OFF          // 1
	FULL_100_MINUS_20        // 2
)

var m = map[int]Pricer{
	ORIGINAL_PRICE:    newNormal(),
	TEN_PERCENT_OFF:   newDiscount(0.9),
	FULL_100_MINUS_20: newReduction(100, 20),
}

// 打折接口
type Pricer interface {
	Price(cash float64) float64
}

// 原价
type normal struct {
}

func newNormal() *normal {
	return &normal{}
}

// 均实现了pricer接口
func (n *normal) Price(cash float64) float64 {
	return cash
}

// 打折
type discount struct {
	percent float64
}

func newDiscount(percent float64) *discount {
	return &discount{percent: percent}
}

func (d *discount) Price(cash float64) float64 {
	return cash * d.percent
}

// 每满fullMoney就减discountMoney
type reduction struct {
	fullMoney     float64
	discountMoney float64
}

func newReduction(fullMoney float64, discountMoney float64) *reduction {
	return &reduction{fullMoney: fullMoney, discountMoney: discountMoney}
}

func (r *reduction) Price(cash float64) float64 {
	count := cash / r.fullMoney
	return cash - count*r.discountMoney
}
