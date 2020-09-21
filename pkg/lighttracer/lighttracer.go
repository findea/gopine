package lighttracer

import (
	"github.com/opentracing/opentracing-go"
	"goweb/pkg/lighttracer/tags"
	"os"
	"strings"
)

type TracerSDKEngine string

const (
	TracerSDKName         = "lighttracer-go"
	TracerSDKVersion      = "1.0.0"
	TracerSDKEngineZipkin = TracerSDKEngine("zipkin")
	TracerSDKEngineJaeger = TracerSDKEngine("jaeger")
)

const (
	OperationTypeHTTP    = "HTTP"
	OperationTypeRPC     = "RPC"
	OperationTypeDB      = "DB"
	OperationTypeTest    = "TEST"
	OperationTypeUnknown = "UNKNOWN"
)

type EndPoint struct {
	ServiceName    string
	ServiceVersion string
	ServiceIpV4    string
	ServiceIpV6    string
	ServicePort    int
}

type Tracer struct {
	OpenTracer      opentracing.Tracer
	TracerSDKEngine TracerSDKEngine
	EndPoint
}

var globalLightTracer Tracer
var _ = opentracing.Tracer(GlobalTracer())

func GlobalTracer() *Tracer {
	return &globalLightTracer
}

func SetGlobalTracer(tracer opentracing.Tracer, localEndPoint EndPoint, sdkEngine TracerSDKEngine) {
	globalLightTracer.OpenTracer = tracer
	globalLightTracer.EndPoint = localEndPoint
	globalLightTracer.TracerSDKEngine = sdkEngine
	opentracing.SetGlobalTracer(tracer)
}

func (l *Tracer) getTracer() opentracing.Tracer {
	if l.OpenTracer != nil {
		return l.OpenTracer
	}

	return opentracing.GlobalTracer()
}

func (l *Tracer) StartSpanWithType(operationName, operationType string, opts ...opentracing.StartSpanOption) opentracing.Span {
	span := l.getTracer().StartSpan(operationName, opts...)
	tags.OperationType.Set(span, operationType)
	tags.Pid.Set(span, os.Getpid())
	tags.TracerName.Set(span, TracerSDKName)
	tags.TracerVersion.Set(span, TracerSDKVersion)
	if l.TracerSDKEngine != "" {
		tags.TracerEngine.Set(span, string(l.TracerSDKEngine))
	}
	if l.ServiceVersion != "" {
		tags.ServiceVersion.Set(span, l.ServiceVersion)
	}

	return span
}

func (l *Tracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return l.StartSpanWithType(operationName, OperationTypeUnknown, opts...)
}

func (l *Tracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return l.getTracer().Inject(sm, format, carrier)
}

func (l *Tracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return l.getTracer().Extract(format, carrier)
}

func (l *Tracer) TraceID(span opentracing.Span) string {
	carrier := make(map[string][]string)
	l.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(carrier))

	items, ok := carrier["X-B3-Traceid"]
	if ok && len(items) > 0 {
		return items[0]
	}

	items, ok = carrier["Uber-Trace-Id"]
	if ok && len(items) > 0 {
		return strings.Split(items[0], ":")[0]
	}

	return ""
}
