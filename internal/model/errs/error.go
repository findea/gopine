package errs

import (
	"goweb/pkg/errors"
	"goweb/pkg/log"
)

var (
	// BadRequest
	BadRequestError        = Errorf(BadRequestErrorCode, "请求错误")
	RequestJsonFormatError = Errorf(BadRequestErrorCode, "JSON格式错误")

	// Unauthorized
	SignTimeStampError     = Errorf(UnauthorizedErrorCode, "时间戳需要当前时间,单位:秒")
	RequestDuplicatedError = Errorf(UnauthorizedErrorCode, "请求重复")
	AppKeyWrongError       = Errorf(UnauthorizedErrorCode, "AppKey错误")
	SignatureError         = Errorf(UnauthorizedErrorCode, "签名错误")
	AppKeyOrTokenError     = Errorf(UnauthorizedErrorCode, "AppKey或Token错误")
	GenTokenFailedError    = Errorf(UnauthorizedErrorCode, "生成Token失败")
)

func ServerError(e error) errors.ErrorWithCode {
	log.Error(e)
	return errors.Error(ServerErrorCode)
}

func BadRequestErrorf(format string, args ...interface{}) errors.ErrorWithCode {
	return errors.Errorf(BadRequestErrorCode, format, args...)
}

func Error(errcode errors.ErrorCode) errors.ErrorWithCode {
	return errors.Error(errcode)
}

func Errorf(errcode errors.ErrorCode, format string, args ...interface{}) errors.ErrorWithCode {
	return errors.Errorf(errcode, format, args...)
}
