package generator

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"strings"
)

type NamingStyle = string

const (
	namingLower NamingStyle = "lower"
	namingCamel NamingStyle = "camel"
	namingSnake NamingStyle = "snake"
)

func formatFilename(filename string, style NamingStyle) string {
	switch style {
	case namingCamel:
		return stringx.From(filename).ToCamel()
	case namingSnake:
		return stringx.From(filename).ToSnake()
	default:
		return strings.ToLower(stringx.From(filename).ToCamel())
	}
}
