package main

import "fmt"

// 砍价购买方式
type BargainPayMethods int32

const (
	BargainPayMethods_BargainPayMethodsAny         BargainPayMethods = 0
	BargainPayMethods_BargainPayMethodsOriginalPay BargainPayMethods = 1 // 原价购买
	BargainPayMethods_BargainPayMethodsDiscountPay BargainPayMethods = 2 // 优惠购买
	BargainPayMethods_BargainPayMethodsAnonPay     BargainPayMethods = 3 // 立即购买
	BargainPayMethods_BargainPayMethodsPay         BargainPayMethods = 4 // 已购买
)

func main()  {
	fmt.Println()
}
