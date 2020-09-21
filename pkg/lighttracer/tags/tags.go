package tags

var (
	ErrNo   = int64TagName("errno")
	ErrStr  = stringTagName("errstr")
	WarnStr = stringTagName("warnstr")
	Version = stringTagName("version")
	Auth    = stringTagName("auth")
	Pid     = intTagName("pid")

	// operation
	OperationType = stringTagName("operation.type")

	// http
	HTTPRequestHeader = stringTagName("http.request.header")
	HTTPRequestBody   = stringTagName("http.request.body")
	HTTPResponseBody  = stringTagName("http.response.body")

	// rpc
	RPCRequestHeader = stringTagName("rpc.request.header")
	RPCRequestBody   = stringTagName("rpc.request.body")
	RPCResponseBody  = stringTagName("rpc.response.body")

	// tracer
	TracerName    = stringTagName("trace.name")
	TracerVersion = stringTagName("trace.version")
	TracerEngine  = stringTagName("trace.engine")

	// service
	ServiceVersion = stringTagName("service.version")

	// db
	DBArgs = stringTagName("db.args")
)
