package uppertrace

import (
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"goweb/pkg/lighttracer"
	"goweb/pkg/lighttracer/tags"
	"upper.io/db.v3"
)

const SpanName = "db.query"

type DBTracerLogger struct {
}

func NewDBTracerLogger() *DBTracerLogger {
	return &DBTracerLogger{}
}

func (lg *DBTracerLogger) Log(m *db.QueryStatus) {
	tracer := lighttracer.GlobalTracer()

	opts := []opentracing.StartSpanOption{
		opentracing.StartTime(m.Start),
	}

	parentSpan := lighttracer.SpanFromContext(m.Context)
	if parentSpan != nil {
		opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
	} else {
		return // 没有，暂时不追踪
	}

	span := tracer.StartSpanWithType(
		SpanName,
		lighttracer.OperationTypeDB,
		opts...,
	)

	ext.DBStatement.Set(span, m.Query)
	ext.DBType.Set(span, "sql")
	ext.Component.Set(span, "upper.io/db.v3")
	if m.RowsAffected != nil {
		span.SetTag("db.rows_affected", *m.RowsAffected)
	}

	bs, err := json.Marshal(m.Args)
	if err == nil {
		tags.DBArgs.Set(span, string(bs))
	}

	if m.Err != nil && m.Err != db.ErrNoMoreRows {
		tags.Error.Set(span, tags.DefaultErrorNo, m.Err.Error())
	}

	span.FinishWithOptions(opentracing.FinishOptions{
		FinishTime: m.End,
	})
}

var _ = db.Logger(NewDBTracerLogger())
