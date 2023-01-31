package main

import (
	"context"
)

type TokenResp struct {
	Token       string
	CreateTime  int64
	IsNewDevice bool
}

type DeviceTokenReq struct {
	Idfa              string
	CaId              string
	Imei              string
	OaId              string
	AndroidId         string
	Os                int
	Token             string
	Source            int
	OldUserCreateTime int
}

type DeviceTokenStrategy interface {
	IsSupport() bool
	GetDeviceToken() (TokenResp, error)
}

type Factory struct {
	ctx context.Context
	req DeviceTokenReq
}

func NewFactory(ctx context.Context, req DeviceTokenReq) *Factory {
	return &Factory{ctx: ctx, req: req}
}

func (f *Factory) GetStrategy() DeviceTokenStrategy {
	var DeviceList []DeviceTokenStrategy
	ctx := f.ctx
	req := f.req
	DeviceList = append(DeviceList, NewAllDeviceIdNotExist(ctx, req))
	DeviceList = append(DeviceList, NewAllDeviceIdNotExist2(ctx, req))
	for _, v := range DeviceList {
		if v.IsSupport() {
			return v
		}
	}
	return nil
}
