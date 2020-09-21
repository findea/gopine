package tags

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type warn struct{}

type err struct{}

var (
	Warn  warn
	Error err
)

const (
	DefaultErrorNo = 1
)

func (err) Set(span opentracing.Span, errno int64, errstr string) {
	ErrNo.Set(span, errno)
	ErrStr.Set(span, errstr)
	if errno != 0 {
		ext.Error.Set(span, true)
	}
}

func (warn) Set(span opentracing.Span, warnstr string) {
	WarnStr.Set(span, warnstr)
}
