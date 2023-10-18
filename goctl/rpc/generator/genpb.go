package generator

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"path/filepath"
	"strings"
)

func (g *Generator) GenPb(ctx DirContext, protoImportPath []string, proto parser.Proto, namingStyle NamingStyle) error {
	dir := ctx.GetPb()
	cw := new(bytes.Buffer)
	base := filepath.Dir(proto.Src)
	cw.WriteString("protoc ")
	for _, ip := range protoImportPath {
		cw.WriteString(" -I=" + ip)
	}
	cw.WriteString(" -I=" + base)
	cw.WriteString(" " + proto.Name)
	if strings.Contains(proto.GoPackage, "/") {
		cw.WriteString(" --go_out=plugins=grpc:" + ctx.GetMain().Filename)
	} else {
		cw.WriteString(" --go_out=plugins=grpc:" + dir.Filename)
	}
	command := cw.String()
	fmt.Println("command:", command)
	g.log.Debug(command)
	_, err := execx.Run(command, "")
	return err
}
