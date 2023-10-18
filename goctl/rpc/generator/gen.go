package generator

import (
	zgenerator "github.com/zeromicro/go-zero/tools/goctl/rpc/generator"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path/filepath"
)

func (g *Generator) Generate(zctx *zgenerator.ZRpcContext) error {
	abs, err := filepath.Abs(zctx.Output)
	if err != nil {
		return err
	}

	err = pathx.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	err = g.ZGenerator.Prepare()
	if err != nil {
		return err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return err
	}

	p := parser.NewDefaultProtoParser()
	proto, err := p.Parse(zctx.Src, zctx.Multiple)
	if err != nil {
		return err
	}

	dirCtx, err := mkdir(projectCtx, proto, g.cfg, zctx)
	if err != nil {
		return err
	}

	//err = g.ZGenerator.GenPb(dirCtx, zctx)
	//if err != nil {
	//	return err
	//}
	//
	err = g.GenLogic(dirCtx, proto, g.cfg, zctx)
	if err != nil {
		return err
	}

	//if zctx.IsGenClient {
	//	err = g.GenCall(dirCtx, proto, g.cfg, zctx)
	//}

	console.NewColorConsole().MarkDone()

	return err
}

// the source file and target location of the rpc service that needs to be generated
//func (g *Generator) Generate(zctx ZRpcContext) error {
//zctx := &ZRpcContext{
//Src: src,
//ProtocCmd: fmt.Sprintf("protoc -I=%s test.proto --go_out=%s --go_opt=Mbase/common.proto=./base --go-grpc_out=%s",
//	common, projectDir, projectDir),
//IsGooglePlugin: true,
//GoOutput:       projectDir,
//GrpcOutput:     projectDir,
//Output:         projectDir,
//}

//abs, err := filepath.Abs(zctx.Output)
//if err != nil {
//	return err
//}
//
//err = util.MkdirIfNotExist(abs)
//if err != nil {
//	return err
//}
//
//projectCtx, err := ctx.Prepare(abs)
//if err != nil {
//	return err
//}
//
////err = g.Generate(ctx)
//
//p := parser.NewDefaultProtoParser()
//proto, err := p.Parse(zctx.Src)
//
//dirCtx, err := mkdir(projectCtx, proto)
//if err != nil {
//	return err
//}

//fmt.Println(proto, err)
//fmt.Println(proto.Name)
//fmt.Println(proto.GoPackage)

//err = g.GenPb(dirCtx, []string{"."}, proto, namingCamel)
//if err != nil {
//	panic(err)
//}
//return g.GenLogic(dirCtx, proto)

//abs, err := filepath.Abs(zctx.Output)
//if err != nil {
//	return err
//}
//
//err = pathx.MkdirIfNotExist(abs)
//if err != nil {
//	return err
//}
//
//err = g.Prepare()
//if err != nil {
//	return err
//}
//
//projectCtx, err := ctx.Prepare(abs)
//if err != nil {
//	return err
//}
//
//p := parser.NewDefaultProtoParser()
//}

//func (g *Generator) genLogicFunction(goPackage string, rpc *parser.RPC) (string, error) {
//	var functions = make([]string, 0)
//	text, err := util.LoadTemplate(category, logicFuncTemplateFileFile, logicFunctionTemplate)
//	if err != nil {
//		return "", err
//	}
//
//	logicName := stringx.From(rpc.Name + "_logic").ToCamel()
//	comment := parser.GetComment(rpc.Doc())
//	buffer, err := util.With("fun").Parse(text).Execute(map[string]interface{}{
//		"logicName":    logicName,
//		"method":       parser.CamelCase(rpc.Name),
//		"request":      fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
//		"response":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
//		"responseType": fmt.Sprintf("%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
//		"hasComment":   len(comment) > 0,
//		"comment":      comment,
//	})
//	if err != nil {
//		return "", err
//	}
//
//	functions = append(functions, buffer.String())
//	return strings.Join(functions, util.NL), nil
//}
