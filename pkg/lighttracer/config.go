package lighttracer

import (
	"fmt"
	zkOpenTracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
	zkLog "github.com/openzipkin/zipkin-go/reporter/log"
	"goweb/pkg/config"
	"log"
	"os"
)

type Config struct {
	ServiceName   string
	ServiceHost   string
	ReporterType  string
	ReporterUrl   string
	BatchSize     int
	BatchInterval config.Duration
}

type Reporter interface {
	Send(model.SpanModel) // Send Span data to the reporter
	Close() error         // Close the reporter
}

func Init(c *Config) error {
	// reporter
	var zkReporter reporter.Reporter

	switch c.ReporterType {
	case "log":
		zkReporter = zkLog.NewReporter(log.New(os.Stdout, "lighttracer", 0))
	case "file":
		fileReporter, err := NewFileReporter(c.ReporterUrl)
		if err != nil {
			return err
		}
		zkReporter = fileReporter
	case "http":
		options := make([]http.ReporterOption, 0)
		if c.BatchSize > 0 {
			options = append(options, http.BatchSize(c.BatchSize))
		}

		if c.BatchInterval.Duration > 0 {
			options = append(options, http.BatchInterval(c.BatchInterval.Duration))
		}

		zkReporter = http.NewReporter(c.ReporterUrl, options...)
	default:
		return fmt.Errorf("%s is not supported", c.ReporterType)
	}

	// endpoint
	if c.ServiceHost == "" {
		hostName, _ := os.Hostname()
		c.ServiceHost = hostName
	}

	endpoint := EndPoint{
		ServiceName: c.ServiceName,
		ServiceIpV4: c.ServiceHost,
	}

	zkEndpoint, err := zipkin.NewEndpoint(c.ServiceName, c.ServiceHost)
	if err != nil {
		return fmt.Errorf("unable to create local endpoint: %+v\n", err)
	}

	// opentracing
	zkTracer, err := zipkin.NewTracer(zkReporter,
		zipkin.WithLocalEndpoint(zkEndpoint),
		zipkin.WithIDGenerator(idgenerator.NewRandomTimestamped()))
	if err != nil {
		return fmt.Errorf("unable to create tracer: %+v\n", err)
	}

	openTracer := zkOpenTracing.Wrap(zkTracer, zkOpenTracing.WithB3InjectOption(zkOpenTracing.B3InjectStandard))

	SetGlobalTracer(openTracer, endpoint, TracerSDKEngineZipkin)

	return nil
}
