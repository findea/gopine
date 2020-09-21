package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/internal/web/common"
	"goweb/pkg/lighttracer"
	gintrace "goweb/pkg/lighttracer/gin"
	"goweb/pkg/log"
	"time"
)

func LoggerHandler(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)
	if raw != "" {
		path = path + "?" + raw
	}

	fileds := map[string]interface{}{
		"StatusCode": c.Writer.Status(),
		"IP":         c.ClientIP(),
		"TimeSpan":   fmt.Sprintf("%.3fs", latency.Seconds()),
		"ErrCode":    c.GetInt64(common.ContextErrCode),
	}

	if span, ok := gintrace.SpanFromContext(c); ok {
		tracer := lighttracer.GlobalTracer()
		traceID := tracer.TraceID(span)
		if traceID != "" {
			fileds["TraceID"] = traceID
		}
	}

	log.WithFields(fileds).Infof("%s %s", method, path)
}
