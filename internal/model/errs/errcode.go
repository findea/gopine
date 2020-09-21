package errs

import "goweb/pkg/errors"

var (
	SuccessCode           = errors.NewErrorCode(0, "请求成功")
	BadRequestErrorCode   = errors.NewErrorCode(400, "请求错误")
	UnauthorizedErrorCode = errors.NewErrorCode(401, "认证错误")
	ServerErrorCode       = errors.NewErrorCode(500, "服务器错误")

	// 注册相关
	EmailAlreadyRegisteredErrorCode = errors.NewErrorCode(100001, "邮箱已经被注册")
	EmailCodeErrorCode              = errors.NewErrorCode(100002, "验证码错误")

	// 登录相关
	EmailNotRegisteredErrorCode   = errors.NewErrorCode(100101, "邮箱未被注册")
	NameOrPassNotMatchedErrorCode = errors.NewErrorCode(100102, "用户名或密码不正确")
)
