package logic

import (
	"context"
	"fmt"

	"pingguoxueyuan/frame/rpc/hello/hello"
	"pingguoxueyuan/frame/rpc/hello/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *hello.Request) (*hello.Response, error) {
	// todo: add your logic here and delete this line

	fmt.Println("this id ping logic")
	return &hello.Response{}, nil
}
