package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/internal/domain"
	"goweb/internal/model/errs"
	"goweb/internal/web/common"
	"time"
)

func NewCounterHandler(max int64, duration time.Duration) gin.HandlerFunc {
	var handler = func(gc *gin.Context) {
		key := fmt.Sprintf("ip:%s", gc.ClientIP())
		count, err := domain.CounterDomain.Count(gc, key, gc.Request.RequestURI, duration)
		if err != nil {
			return
		}

		if count > max {
			common.ResponseErr(gc, errs.BadRequestErrorf(errs.RequestTooManyErrorStr))
			gc.Abort()
		}
	}

	return handler
}
