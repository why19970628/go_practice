package parser

import (
	"github.com/emicklei/proto"
)

// Proto describes a proto file,
type Proto struct {
	Src       string
	Name      string
	Package   Package
	PbPackage string
	GoPackage string
	Import    []Import
	Message   []Message
	Service   Services
}

type (
	// Services is a slice of Service.
	Services []Service

	// Service describes the rpc service, which is the relevant
	// content after the translation of the proto file
	Service struct {
		*proto.Service
		RPC []*RPC
	}
)

// Package defines the protobuf package.
type Package struct {
	*proto.Package
}

// Message embeds proto.Message
type Message struct {
	*proto.Message
}

// Import embeds proto.Import
type Import struct {
	*proto.Import
}

// RPC embeds proto.RPC
type RPC struct {
	*proto.RPC
}
