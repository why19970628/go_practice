package main

import (
	"context"
	"fmt"
)

func main() {
	strategy := NewFactory(context.Background(), DeviceTokenReq{}).GetStrategy()

	resp, _ := strategy.GetDeviceToken()
	fmt.Println(fmt.Sprintf("%+v", resp))
}
