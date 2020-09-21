package domain

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"goweb/internal/conf"
	"goweb/internal/model/auth"
	"goweb/internal/model/errs"
	"goweb/pkg/util/json"
	"strings"
	"time"
)

type jwtDomain struct {
}

var JWTDomain = new(jwtDomain)

func (d *jwtDomain) GenerateJWTToken(c context.Context, claims *auth.UserClaims) (string, error) {
	return d.CreateToken(c, conf.Conf.JWT.SecretBytes(), claims, time.Now().Add(conf.Conf.JWT.ExpiresAt.Duration).Unix())
}

func (d *jwtDomain) VerifyJWTToken(c context.Context, token string) (*auth.UserClaims, error) {
	return d.ParseToken(c, conf.Conf.JWT.SecretBytes(), token)
}

func (d *jwtDomain) CreateToken(_ context.Context, secretKey []byte, claims *auth.UserClaims, expiresAt int64) (tokenString string, err error) {
	var jwtCustomClaims struct {
		auth.UserClaims
		jwt.StandardClaims
	}

	jwtCustomClaims.UserClaims = *claims
	jwtCustomClaims.ExpiresAt = expiresAt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims)
	tokenString, err = token.SignedString(secretKey)
	return
}

func (d *jwtDomain) ParseToken(_ context.Context, secretKey []byte, tokenString string) (claims *auth.UserClaims, err error) {
	defer func() {
		if err == nil {
			return
		}

		validationError, ok := err.(*jwt.ValidationError)
		switch {
		case ok && validationError.Errors&jwt.ValidationErrorExpired > 0:
			err = errs.Errorf(errs.UnauthorizedErrorCode, errs.TokenExpiredErrorStr)
		case ok && validationError.Errors&jwt.ValidationErrorMalformed > 0:
			err = errs.Errorf(errs.UnauthorizedErrorCode, errs.TokenMalformedErrorStr)
		default:
			err = errs.ServerError(err)
		}

	}()

	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return
	}

	if err = token.Claims.Valid(); err != nil {
		return
	}

	claimsBase64Str := strings.Split(tokenString, ".")[1]
	claimsBytes, err := jwt.DecodeSegment(claimsBase64Str)
	if err != nil {
		return nil, err
	}

	claims = new(auth.UserClaims)
	err = json.FromJsonBytes(claimsBytes, claims)
	if err != nil {
		return nil, err
	}

	return
}
