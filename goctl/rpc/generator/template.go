package generator

const (
	category                          = "rpc"
	callTemplateFile                  = "call.tpl"
	callInterfaceFunctionTemplateFile = "call-interface-func.tpl"
	callFunctionTemplateFile          = "call-func.tpl"
	configTemplateFileFile            = "config.tpl"
	etcTemplateFileFile               = "etc.tpl"
	logicTemplateFileFile             = "logic.tpl"
	logicFuncTemplateFileFile         = "logic-func.tpl"
	mainTemplateFile                  = "main.tpl"
	serverTemplateFile                = "server.tpl"
	serverFuncTemplateFile            = "server-func.tpl"
	svcTemplateFile                   = "svc.tpl"
	rpcTemplateFile                   = "template.tpl"
)

var templates = map[string]string{
	//callTemplateFile:                  callTemplateText,
	callInterfaceFunctionTemplateFile: callInterfaceFunctionTemplate,
	callFunctionTemplateFile:          callFunctionTemplate,
	//configTemplateFileFile:            configTemplate,
	//etcTemplateFileFile:               etcTemplate,
	//logicTemplateFileFile:             logicTemplate,
	//logicFuncTemplateFileFile:         logicFunctionTemplate,
	//mainTemplateFile:                  mainTemplate,
	//serverTemplateFile:                serverTemplate,
	//serverFuncTemplateFile:            functionTemplate,
	//svcTemplateFile:                   svcTemplate,
	//rpcTemplateFile:                   rpcTemplateText,
}
