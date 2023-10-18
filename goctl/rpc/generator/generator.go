package generator

import (
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	zgenerator "github.com/zeromicro/go-zero/tools/goctl/rpc/generator"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"log"
)

type Generator struct {
	log        console.Console
	cfg        *conf.Config
	ZGenerator *zgenerator.Generator
}

func (g *Generator) GetCfg() *conf.Config {
	return g.cfg
}

// NewGenerator returns an instance of Generator
func NewGenerator(style string, verbose bool) *Generator {
	cfg, err := conf.NewConfig(style)
	if err != nil {
		log.Fatalln(err)
	}

	colorLogger := console.NewColorConsole(verbose)

	return &Generator{
		log:        colorLogger,
		cfg:        cfg,
		ZGenerator: zgenerator.NewGenerator(style, verbose),
	}
}
