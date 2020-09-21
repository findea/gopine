package di

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"goweb/pkg/lighttracer"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	tracer := lighttracer.GlobalTracer()

	span := tracer.StartSpanWithType("trace_test", lighttracer.OperationTypeTest)

	time.Sleep(10 * time.Millisecond)
	span.SetTag("tag_key", "tag_value")

	dbSpan := tracer.StartSpanWithType("dbquery", lighttracer.OperationTypeDB, opentracing.ChildOf(span.Context()))
	ext.DBStatement.Set(dbSpan, "select * from user where user_id = ?")
	time.Sleep(10 * time.Millisecond)
	dbSpan.Finish()

	time.Sleep(10 * time.Millisecond)
	span.LogKV("event_key", "event_value")

	time.Sleep(10 * time.Millisecond)
	span.Finish()

	time.Sleep(10 * time.Millisecond)
	carrier := make(map[string][]string)
	tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(carrier))
	t.Logf("%+v", carrier)
}
