package main

import "context"

type AllDeviceIdNotExist struct {
	ctx context.Context
	req DeviceTokenReq
}

func NewAllDeviceIdNotExist(ctx context.Context, req DeviceTokenReq) *AllDeviceIdNotExist {
	return &AllDeviceIdNotExist{ctx: ctx, req: req}
}

func (a *AllDeviceIdNotExist) IsSupport() bool {
	return false
}

func (a *AllDeviceIdNotExist) GetDeviceToken() (TokenResp, error) {
	return TokenResp{
		Token: "123",
	}, nil
}

type AllDeviceIdNotExist2 struct {
	ctx context.Context
	req DeviceTokenReq
}

func NewAllDeviceIdNotExist2(ctx context.Context, req DeviceTokenReq) *AllDeviceIdNotExist2 {
	return &AllDeviceIdNotExist2{ctx: ctx, req: req}
}

func (a *AllDeviceIdNotExist2) IsSupport() bool {
	return true
}

func (a *AllDeviceIdNotExist2) GetDeviceToken() (TokenResp, error) {
	return TokenResp{
		Token: "456",
	}, nil
}
