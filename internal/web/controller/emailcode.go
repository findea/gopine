package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/model/request"
	"goweb/internal/service"
	"goweb/internal/web/common"
)

type emailCodeController struct {
}

var EmailCodeController = new(emailCodeController)

func (c *emailCodeController) Send(gc *gin.Context) {
	reqObj := new(request.SendEmailCodeReq)
	err := common.ParseReqFromBody(gc, reqObj)
	if err != nil {
		common.ResponseErr(gc, err)
		return
	}

	resp, err := service.EmailCodeService.SendEmailToken(gc, reqObj)
	common.Response(gc, resp, err)
}
