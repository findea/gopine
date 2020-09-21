package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/domain"
	"goweb/internal/model/auth"
	"goweb/internal/model/errs"
	"goweb/internal/web/common"
	"net/http"
	"strings"
)

const GinContextJWTKey = "GinContextJWTKey"
const JWTTokenCookieName = "token"

func JWTTokenVerifyHandler(c *gin.Context) {
	token, err := c.Cookie(JWTTokenCookieName)
	if err != nil {
		token, err = getJWTTokenFromAuthorization(c)
		if err != nil {
			common.ResponseErr(c, err)
			c.Abort()
			return
		}
	}

	claims, err := domain.JWTDomain.VerifyJWTToken(c, token)
	if err != nil {
		common.ResponseErr(c, err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set(GinContextJWTKey, claims)
	c.Next()
}

func getJWTTokenFromAuthorization(c *gin.Context) (string, error) {
	const bearerLength = len("Bearer ")
	authorization := c.GetHeader("Authorization")
	if len(authorization) < bearerLength {
		return "", errs.Errorf(errs.UnauthorizedErrorCode, errs.TokenMalformedErrorStr)
	}
	return strings.TrimSpace(authorization[bearerLength:]), nil
}

func GetUserClaims(c *gin.Context) (*auth.UserClaims, error) {
	val, exists := c.Get(GinContextJWTKey)
	if !exists {
		return nil, errs.Error(errs.UnauthorizedErrorCode)
	}

	claims, ok := val.(*auth.UserClaims)
	if !ok {
		return nil, errs.Error(errs.UnauthorizedErrorCode)
	}

	return claims, nil
}
