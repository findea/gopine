package service

import (
	"context"
	"goweb/internal/domain"
	"goweb/internal/model/errs"
	"goweb/internal/model/request"
	"goweb/pkg/util/validate"
)

type emailCodeService struct {
}

var EmailCodeService = new(emailCodeService)

func (s *emailCodeService) SendEmailToken(c context.Context, req *request.SendEmailCodeReq) (*request.SendEmailCodeResp, error) {
	errMsg := validate.ValidateObject(req)
	if errMsg != "" {
		return nil, errs.BadRequestErrorf(errMsg)
	}

	code := domain.EmailCodeDomain.Rand4DigistCode(c)
	err := domain.EmailCodeDomain.SendEmailCode(c, req.Email, code)
	if err != nil {
		return nil, err
	}

	return &request.SendEmailCodeResp{}, nil
}
