package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFuncTimeOut(t *testing.T) {
	type args struct {
		ctx     context.Context
		timeOut time.Duration
		fn      func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "测试超时",
			args: args{
				ctx:     context.Background(),
				timeOut: time.Second,
				fn: func() error {
					time.Sleep(time.Second * 2)
					return nil
				},
			},
			wantErr: context.DeadlineExceeded,
		},
		{
			name: "测试函数报错",
			args: args{
				ctx:     context.Background(),
				timeOut: time.Second,
				fn: func() error {
					return errors.New("函数报错")
				},
			},
			wantErr: errors.New("函数报错"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FuncTimeOut(tt.args.ctx, tt.args.timeOut, tt.args.fn); err.Error() != tt.wantErr.Error() {
				t.Errorf("FuncTimeOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
