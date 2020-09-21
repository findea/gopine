package lighttracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func userLogin(t *testing.T, tracer Tracer) {
	span := tracer.StartSpanWithType("/user/login", OperationTypeHTTP)
	assert.NotNil(t, span)
	defer span.Finish()
	time.Sleep(time.Second)
}

func rpcRequest(t *testing.T, clientTracer Tracer, serverTracer Tracer) {
	// client
	clientSpan := clientTracer.StartSpanWithType(
		"/api/sum",
		OperationTypeHTTP,
		ext.SpanKindRPCClient,
	)
	defer clientSpan.Finish()

	header := make(map[string][]string)
	err := clientTracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	assert.Nil(t, err)
	fmt.Println(header)

	// server
	spanCtx, err := serverTracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	assert.Nil(t, err)

	serverSpan := serverTracer.StartSpanWithType(
		"/api/sum",
		OperationTypeHTTP,
		ext.SpanKindRPCServer,
		opentracing.ChildOf(spanCtx),
	)
	defer serverSpan.Finish()

	time.Sleep(time.Second)
}

func makeLocalEndPoint() EndPoint {
	return EndPoint{
		ServiceName:    "user_serive",
		ServiceVersion: "1.0.0",
		ServiceIpV4:    "22.33.44.55",
		ServicePort:    80,
	}
}

func makeRPCClientEndPoint() EndPoint {
	return EndPoint{
		ServiceName:    "rpc_client",
		ServiceVersion: "1.0.0",
		ServiceIpV4:    "22.33.44.11",
		ServicePort:    80,
	}
}

func makeRPCServerEndPoint() EndPoint {
	return EndPoint{
		ServiceName:    "rpc_service",
		ServiceVersion: "1.0.0",
		ServiceIpV4:    "22.33.44.22",
		ServicePort:    80,
	}
}
