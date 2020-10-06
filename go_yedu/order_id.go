package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	return fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func GetDateTime() string {
	t := Current()
	return t.Format(TIME_INT_LAYOUT)
}

func WithDrawOrderId(uid string, preStr string) string {
	return preStr + GetDateTime() + uid[0:4] + CreateCaptcha()
}

func GetWithDrawOrderNo(uid string) string {
	if len(uid) < 5 {
		return GetCurrentDateTimeStr() + uid + CreateCaptcha()
	}
	return GetCurrentDateTimeStr() + uid[:5] + CreateCaptcha()
}



func main() {
	fmt.Println(CreateCaptcha())
}

