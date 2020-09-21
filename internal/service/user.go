package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"goweb/internal/domain"
	"goweb/internal/model/auth"
	"goweb/internal/model/dbmodel"
	"goweb/internal/model/errs"
	"goweb/internal/model/request"
	"goweb/pkg/log"
	"goweb/pkg/util/validate"
)

type userService struct {
}

var UserService = new(userService)

func (s *userService) Register(c context.Context, req *request.UserRegisterReq) (*request.UserRegisterResp, error) {
	log.Debugf("add user with %+v", req)

	// 参数检查
	errMsg := validate.ValidateObject(req)
	if errMsg != "" {
		return nil, errs.BadRequestErrorf(errMsg)
	}

	// 检查邮箱是否已注册
	userInDB, err := domain.UserDomain.SelectOneByEmail(c, req.Email)
	if err != nil {
		return nil, errs.ServerError(err)
	}
	if userInDB != nil {
		return nil, errs.Errorf(errs.EmailAlreadyRegisteredErrorCode, errs.EmailAlreadyRegistedErrStr, req.Email)
	}

	// 检查 验证码
	ok, err := domain.EmailCodeDomain.VerifyEmailCode(c, req.Email, req.Code)
	if err != nil {
		return nil, errs.ServerError(err)
	}
	if !ok {
		return nil, errs.Error(errs.EmailCodeErrorCode)
	}

	// 密码加密存储
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.ServerError(err)
	}
	var password = string(hash)

	// 插入数据库
	var user = &dbmodel.User{
		Nickname:    req.Nickname,
		Email:       req.Email,
		Password:    password,
		LastLoginIP: req.LastLoginIP,
	}
	_, err = domain.UserDomain.Insert(c, user)
	if err != nil {
		return nil, errs.ServerError(err)
	}

	return &request.UserRegisterResp{}, nil
}

func (s *userService) Login(c context.Context, req *request.UserLoginReq) (*request.UserLoginResp, error) {
	log.Debugf("user login with %+v", req)

	// 参数检查
	errMsg := validate.ValidateObject(req)
	if errMsg != "" {
		return nil, errs.BadRequestErrorf(errMsg)
	}

	// 检查邮箱是否已注册
	userInDB, err := domain.UserDomain.SelectOneByEmail(c, req.Email)
	if err != nil {
		return nil, errs.ServerError(err)
	}
	if userInDB == nil {
		return nil, errs.Errorf(errs.EmailNotRegisteredErrorCode, errs.EmailNotRegistedErrorStr, req.Email)
	}

	// 检查密码是否相等
	err = bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(req.Password))
	if err != nil {
		return nil, errs.Errorf(errs.NameOrPassNotMatchedErrorCode, errs.NameOrPassNotMatchedErrorStr)
	}

	// 生成 JWT
	userCliams := &auth.UserClaims{
		UserID: userInDB.UserID,
		Email:  userInDB.Email,
	}
	token, err := domain.JWTDomain.GenerateJWTToken(c, userCliams)
	if err != nil {
		return nil, errs.ServerError(err)
	}

	return &request.UserLoginResp{
		Nickname: userInDB.Nickname,
		Token:    token,
	}, nil
}
