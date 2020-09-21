package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/model/request"
	"goweb/internal/service"
	"goweb/internal/web/common"
	"goweb/internal/web/middleware"
	"net/http"
)

type userController struct {
}

var UserController = new(userController)

func (c *userController) Register(gc *gin.Context) {
	reqObj := new(request.UserRegisterReq)
	err := common.ParseReqFromBody(gc, reqObj)
	if err != nil {
		common.ResponseErr(gc, err)
		return
	}

	reqObj.LastLoginIP = gc.ClientIP()

	resp, err := service.UserService.Register(gc, reqObj)
	common.Response(gc, resp, err)
}

func (c *userController) Login(gc *gin.Context) {
	reqObj := new(request.UserLoginReq)
	err := common.ParseReqFromBody(gc, reqObj)
	if err != nil {
		common.ResponseErr(gc, err)
		return
	}

	reqObj.LastLoginIP = gc.ClientIP()

	resp, err := service.UserService.Login(gc, reqObj)
	if err == nil {
		cookie := &http.Cookie{
			Name:     middleware.JWTTokenCookieName,
			Value:    resp.Token,
			HttpOnly: true,
		}
		http.SetCookie(gc.Writer, cookie)
	}

	common.Response(gc, resp, err)
}
