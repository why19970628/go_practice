package main

import "fmt"

// 下面是一大堆自定义的Error
type RedisEvalError struct {
}

func (e RedisEvalError) Error() string {
	return "Error when executing redisService eval."
}

type UserHasCouponError struct {
	userName   string
	couponName string
}

func (e UserHasCouponError) Error() string {
	return fmt.Sprintf("User %s has had coupon %s.", e.userName, e.couponName)
}

type NoSuchCouponError struct {
	userName   string
	couponName string
}

func (e NoSuchCouponError) Error() string {
	return fmt.Sprintf("Coupon %s created by %s doesn't exist.", e.couponName, e.userName)
}

type NoCouponLeftError struct {
	userName   string
	couponName string
}

func (e NoCouponLeftError) Error() string {
	return fmt.Sprintf("No Coupon %s created by %s left.", e.couponName, e.userName)
}

type CouponLeftResError struct {
	couponLeftRes interface{}
}

func (e CouponLeftResError) Error() string {
	switch e.couponLeftRes.(type) {
	case int:
		return fmt.Sprintf("Unexpected couponLeftRes Num: %v.", e.couponLeftRes)
	default:
		return fmt.Sprintf("couponLeftRes : %v with wrong type.", e.couponLeftRes)
	}
}
