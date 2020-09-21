package gintrace

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"goweb/pkg/lighttracer"
	"goweb/pkg/lighttracer/tags"
	"io/ioutil"
	"net/http"
	"strings"
)

const contextTracerKey = "LightTracerContext"

func TraceHandler(c *gin.Context) {
	opts := []opentracing.StartSpanOption{
		ext.SpanKindRPCServer,
		opentracing.Tag{Key: string(ext.Component), Value: "gin"},
		opentracing.Tag{Key: string(ext.PeerHostIPv4), Value: c.Request.RemoteAddr},
	}

	tracer := lighttracer.GlobalTracer()
	spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err == nil {
		opts = append(opts, opentracing.ChildOf(spanCtx))
	}

	// start span
	span := tracer.StartSpanWithType(c.Request.RequestURI, lighttracer.OperationTypeHTTP, opts...)
	defer span.Finish()
	ext.HTTPMethod.Set(span, c.Request.Method)
	ext.HTTPUrl.Set(span, c.Request.URL.EscapedPath())

	// header
	if data, err := json.Marshal(c.Request.Header); err == nil {
		tags.HTTPRequestHeader.Set(span, string(data))
	}

	// body
	body := ""
	contentType := c.ContentType()
	if !isBinaryContent(contentType) && !isMultipart(contentType) {
		data, _ := c.GetRawData()
		if data != nil {
			body = string(data)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
	}
	tags.HTTPRequestBody.Set(span, body)

	// context
	setSpanContext(c, span)

	// writer
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	// next
	c.Next()

	// response
	tags.HTTPResponseBody.Set(span, blw.body.String())

	statusCode := c.Writer.Status()
	ext.HTTPStatusCode.Set(span, uint16(statusCode))

	if statusCode >= http.StatusInternalServerError {
		tags.Error.Set(span, int64(statusCode), "server err")
	}
}

func SpanFromContext(c *gin.Context) (opentracing.Span, bool) {
	traceCtx, ok := getSpanContext(c)
	if ok == false {
		return nil, false
	}
	return lighttracer.SpanFromContext(traceCtx), true
}

func setSpanContext(c *gin.Context, span opentracing.Span) {
	ctx := lighttracer.ContextWithSpan(context.TODO(), span)
	c.Set(contextTracerKey, ctx)
}

func getSpanContext(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get(contextTracerKey)
	if exist == false {
		ok = false
		ctx = context.TODO()
		return
	}
	ctx, ok = v.(context.Context)
	return
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func isBinaryContent(contentType string) bool {
	return strings.Contains(contentType, "image") ||
		strings.Contains(contentType, "video") ||
		strings.Contains(contentType, "audio")
}

func isMultipart(contentType string) bool {
	return strings.Contains(contentType, "multipart/form-data")
}
