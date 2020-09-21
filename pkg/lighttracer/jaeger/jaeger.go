package jaeger

import (
	"goweb/pkg/lighttracer"
	"io"
	"time"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func SetGlobalJaegerTracer(servicename string, addr string) (*lighttracer.Tracer, io.Closer, error) {
	// config
	cfg := config.Configuration{
		ServiceName: servicename,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// report
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(sender)

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		config.Reporter(reporter),
	)

	lighttracer.SetGlobalTracer(tracer, lighttracer.EndPoint{
		ServiceName: servicename,
	}, lighttracer.TracerSDKEngineJaeger)

	return lighttracer.GlobalTracer(), closer, err
}
