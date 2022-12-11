package main

// 生成具体策略的工厂
type CashFactory struct {
}

// 将具体的策略实现封装起来，交由工厂来实现
func (c CashFactory) GetCashPrice(discountType int) Pricer {
	h := m[discountType]
	return h

	//var cashPricer Pricer
	//switch discountType {
	//case ORIGINAL_PRICE:
	//	cashPricer = &normal{}
	//case TEN_PERCENT_OFF:
	//	cashPricer = &discount{0.9}
	//case FULL_100_MINUS_20:
	//	cashPricer = &reduction{100, 20}
	//}
	//return cashPricer
}
