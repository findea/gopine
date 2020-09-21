package common

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/model/errs"
	"goweb/pkg/util/json"
	"io/ioutil"
)

func ParseReqFromBody(c *gin.Context, req interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return errs.RequestJsonFormatError
	}

	bodyStr := string(body)
	err = json.FromJson(bodyStr, req)
	if err != nil {
		return errs.RequestJsonFormatError
	}

	return nil
}
