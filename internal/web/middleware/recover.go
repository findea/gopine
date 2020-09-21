package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gintrace "goweb/pkg/lighttracer/gin"
	"goweb/pkg/lighttracer/tags"
	"goweb/pkg/log"
	"net/http"
	"net/http/httputil"
	"runtime"
)

func RecoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			errMsg := fmt.Sprintf("%s\n%s\n%s", string(httprequest), err, buf)
			log.Error(errMsg)

			if span, ok := gintrace.SpanFromContext(c); ok {
				tags.Error.Set(span, http.StatusInternalServerError, errMsg)
			}

			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
