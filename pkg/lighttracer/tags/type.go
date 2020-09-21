package tags

import "github.com/opentracing/opentracing-go"

type stringTagName string

// Set adds a string tag to the `span`
func (tag stringTagName) Set(span opentracing.Span, value string) {
	span.SetTag(string(tag), value)
}

type intTagName string

func (tag intTagName) Set(span opentracing.Span, value int) {
	span.SetTag(string(tag), value)
}

type int32TagName string

// Set adds a int32 tag to the `span`
func (tag int32TagName) Set(span opentracing.Span, value int32) {
	span.SetTag(string(tag), value)
}

type int64TagName string

// Set adds a int64 tag to the `span`
func (tag int64TagName) Set(span opentracing.Span, value int64) {
	span.SetTag(string(tag), value)
}

type boolTagName string

// Register adds a bool tag to the `span`
func (tag boolTagName) Set(span opentracing.Span, value bool) {
	span.SetTag(string(tag), value)
}

type errorTagName string

// Register adds a bool tag to the `span`
func (tag errorTagName) Set(span opentracing.Span, value error) {
	span.SetTag(string(tag), value)
}
