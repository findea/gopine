package errors

import "fmt"

var (
	errcodes = make(map[int64]ErrorCode)
	_        = ErrorWithCode(&ErrorCode{})
)

type ErrorCode struct {
	err  error
	code int64
	desc string
}

func (c *ErrorCode) Error() string {
	if c.err == nil {
		return c.desc
	}
	return c.err.Error()
}

func (c *ErrorCode) Code() int64 {
	return c.code
}

func NewErrorCode(code int64, desc string) ErrorCode {
	if val, ok := errcodes[code]; ok {
		panic(fmt.Sprintf("error code: %d already exist for %v", code, val))
	}

	errorCode := ErrorCode{
		code: code,
		desc: desc,
	}
	errcodes[code] = errorCode
	return errorCode
}
