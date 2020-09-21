package common

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/model/errs"
	"goweb/pkg/errors"
	"goweb/pkg/lighttracer/gin"
	"goweb/pkg/lighttracer/tags"
	"net/http"
)

const (
	ContextErrCode = "context/err/code"
)

type response struct {
	ErrCode int64       `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Result  interface{} `json:"result,omitempty"`
}

func Response(c *gin.Context, data interface{}, err error) {
	if err != nil {
		ResponseErr(c, err)
		return
	}
	ResponseResult(c, data)
}

func ResponseResult(c *gin.Context, result interface{}) {
	resposeWithTrace(c, &response{
		ErrCode: errs.SuccessCode.Code(),
		Result:  result,
	})
}

func ResponseErr(c *gin.Context, e error) {
	errWrapper, ok := e.(errors.ErrorWithCode)
	if ok && errWrapper == nil {
		errWrapper = errs.BadRequestError
	}

	resposeWithTrace(c, &response{
		ErrCode: errWrapper.Code(),
		ErrMsg:  errWrapper.Error(),
	})
}

func ResponseErrStr(c *gin.Context, err error) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusBadRequest, err.Error())
}

func ResponseSvgStr(c *gin.Context, str string) {
	c.Writer.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	c.String(http.StatusOK, str)
}

func resposeWithTrace(c *gin.Context, r *response) {
	c.Set(ContextErrCode, r.ErrCode)
	c.JSON(http.StatusOK, r)

	if span, ok := gintrace.SpanFromContext(c); ok {
		if r.ErrCode == errs.ServerErrorCode.Code() {
			tags.Error.Set(span, int64(r.ErrCode), r.ErrMsg)
		} else if r.ErrCode != errs.SuccessCode.Code() {
			tags.Warn.Set(span, r.ErrMsg)
		}
	}
}
