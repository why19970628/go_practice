package generator

//const (
//	logicTemplate = `package logic
//
//import (
//	"context"
//
//	{{.imports}}
//
//	"github.com/zeromicro/go-zero/core/logx"
//)
//
//type {{.logicName}} struct {
//	ctx    context.Context
//	svcCtx *svc.ServiceContext
//	logx.Logger
//}
//
//func New{{.logicName}}(ctx context.Context,svcCtx *svc.ServiceContext) *{{.logicName}} {
//	return &{{.logicName}}{
//		svcCtx: svcCtx,
//		ctx:    ctx,
//		Logger: logx.WithContext(ctx),
//	}
//}
//{{.functions}}
//`
//	logicFunctionTemplate = `{{if .hasComment}}{{.comment}}{{end}}
//func (l *{{.logicName}}) {{.method}} (in {{.request}}) ({{.response}}, error) {
//	// todo: add your logic here and delete this line
//
//	return &{{.responseType}}{}, nil
//}
//`
//)
