package main

import "fmt"

// https://studygolang.com/articles/5199
//策略模式定义了算法家族，在调用算法家族的时候不感知算法的变化，客户也不会受到影响。
//
//下面用《大话设计模式》中的一个实例进行改写。
//
//例：超市中经常进行促销活动，促销活动的促销方法就是一个个策略，如“满一百减20”，“打八折”等。现在实现策略模式，用CashContext生产策略，并完成策略的调用。

type cashSuper interface {
	AcceptMoney(money float64) float64
}

//普通情况，没有折扣
type cashNormal struct {
}

func newCashNormal() cashNormal {
	instance := new(cashNormal)
	return *instance
}

func (c cashNormal) AcceptMoney(money float64) float64 {
	return money
}

//打折，传入打折的折扣，如0.8
type cashDiscount struct {
	Rebate float64 //折扣
}

func newCashDiscount(rebate float64) cashDiscount {
	instance := new(cashDiscount)
	instance.Rebate = rebate
	return *instance
}

func (c cashDiscount) AcceptMoney(money float64) float64 {
	return money * c.Rebate
}

//直接返利，如满100返20
type cashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func newCashReturn(moneyCondition float64, moneyReturn float64) cashReturn {
	instance := new(cashReturn)
	instance.MoneyCondition = moneyCondition
	instance.MoneyReturn = moneyReturn
	return *instance
}

func (c cashReturn) AcceptMoney(money float64) float64 {
	if money >= c.MoneyCondition {
		moneyMinus := int(money / c.MoneyCondition)
		return money - float64(moneyMinus)*c.MoneyReturn
	}
	return money
}

// 定义CashContext结构，用来做策略筛选
type CashContext struct {
	Strategy cashSuper
}

func NewCashContext(cashType string) CashContext {
	c := new(CashContext)
	//这里事实上是简易工厂模式的变形，用来生产策略
	switch cashType {
	case "打八折":
		c.Strategy = newCashDiscount(0.8)
	case "满一百返20":
		c.Strategy = newCashReturn(100.0, 20.0)
	default:
		c.Strategy = newCashNormal()
	}
	return *c
}

//在策略生产成功后，我们就可以直接调用策略的函数。
func (c CashContext) GetMoney(money float64) float64 {
	return c.Strategy.AcceptMoney(money)
}

func main() {
	money := 100.0
	cc := NewCashContext("打八折")
	money = cc.GetMoney(money)
	fmt.Println("100打八折实际金额为", money)

	money = 199
	cc = NewCashContext("满一百返20")
	money = cc.GetMoney(money)
	fmt.Println("199满一百返20实际金额为", money)

	money = 199
	cc = NewCashContext("没有折扣")
	money = cc.GetMoney(money)
	fmt.Println("199没有折扣实际金额为", money)

}
