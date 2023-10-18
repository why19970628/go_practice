package cli

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/generator"
	mygenerator "goctl/rpc/generator"
)

//
//func RPCNew(cmd *cobra.Command, args []string) error {
//	rpcname := args[0]
//	ext := filepath.Ext(rpcname)
//	if len(ext) > 0 {
//		return fmt.Errorf("unexpected ext: %s", ext)
//	}
//	style := VarStringStyle
//	home := VarStringHome
//	remote := VarStringRemote
//	branch := VarStringBranch
//	verbose := VarBoolVerbose
//	if len(remote) > 0 {
//		repo, _ := util.CloneIntoGitHome(remote, branch)
//		if len(repo) > 0 {
//			home = repo
//		}
//	}
//	if len(home) > 0 {
//		pathx.RegisterGoctlHome(home)
//	}
//
//	protoName := rpcname + ".proto"
//	filename := filepath.Join(".", rpcname, protoName)
//	src, err := filepath.Abs(filename)
//	if err != nil {
//		return err
//	}
//
//	err = generator.ProtoTmpl(src)
//	if err != nil {
//		return err
//	}
//
//	var ctx generator.ZRpcContext
//	ctx.Src = src
//	ctx.GoOutput = filepath.Dir(src)
//	ctx.GrpcOutput = filepath.Dir(src)
//	ctx.IsGooglePlugin = true
//	ctx.Output = filepath.Dir(src)
//	ctx.ProtocCmd = fmt.Sprintf("protoc -I=%s %s --go_out=%s --go-grpc_out=%s", filepath.Dir(src), filepath.Base(src), filepath.Dir(src), filepath.Dir(src))
//	ctx.IsGenClient = VarBoolClient
//
//	grpcOptList := VarStringSliceGoGRPCOpt
//	if len(grpcOptList) > 0 {
//		ctx.ProtocCmd += " --go-grpc_opt=" + strings.Join(grpcOptList, ",")
//	}
//
//	goOptList := VarStringSliceGoOpt
//	if len(goOptList) > 0 {
//		ctx.ProtocCmd += " --go_opt=" + strings.Join(goOptList, ",")
//	}
//
//	g := generator.NewGenerator(style, verbose)
//	return g.Generate(&ctx)
//}

func RPCNew(cmd *cobra.Command, args []string) error {

	ctx := generator.ZRpcContext{
		Src: "./live.proto",
		//ProtocCmd: fmt.Sprintf("protoc -I=%s live.proto --go_out=%s --go_opt=Mbase/common.proto=./base --go-grpc_out=%s",
		//	"./proto/live", "./proto/live", "./proto/live"),
		//IsGooglePlugin: true,
		//GoOutput:       projectDir,
		//GrpcOutput:     projectDir,
		//Output:         projectDir,
	}

	err := mygenerator.NewGenerator("", true).Generate(&ctx)
	if err != nil {
		panic(err)
	}
	return nil
}
