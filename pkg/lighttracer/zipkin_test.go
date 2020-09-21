package lighttracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	zkOpenTracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/idgenerator"
	zkLog "github.com/openzipkin/zipkin-go/reporter/log"
	"log"
	"os"
	"testing"
)

func TestZipkinUserLogin(t *testing.T) {
	localEndPoint := makeLocalEndPoint()
	userLogin(
		t,
		Tracer{
			OpenTracer:      newZipkinTracer(localEndPoint),
			EndPoint:        localEndPoint,
			TracerSDKEngine: TracerSDKEngineZipkin,
		},
	)
}

func TestZipkinRPC(t *testing.T) {
	clientEndPoint := makeRPCClientEndPoint()
	serverEndPoint := makeRPCServerEndPoint()
	rpcRequest(
		t,
		Tracer{
			OpenTracer:      newZipkinTracer(clientEndPoint),
			EndPoint:        clientEndPoint,
			TracerSDKEngine: TracerSDKEngineZipkin,
		},
		Tracer{
			OpenTracer:      newZipkinTracer(serverEndPoint),
			EndPoint:        serverEndPoint,
			TracerSDKEngine: TracerSDKEngineZipkin,
		},
	)
}

func newZipkinTracer(endPoint EndPoint) opentracing.Tracer {
	reporter := zkLog.NewReporter(log.New(os.Stdout, "", 0))
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint(endPoint.ServiceName,
		fmt.Sprintf("%s:%d", endPoint.ServiceIpV4, endPoint.ServicePort))

	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	zkTracer, err := zipkin.NewTracer(reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithIDGenerator(idgenerator.NewRandomTimestamped()))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	return zkOpenTracing.Wrap(zkTracer, zkOpenTracing.WithB3InjectOption(zkOpenTracing.B3InjectStandard))
}
