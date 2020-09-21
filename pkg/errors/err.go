package errors

import "fmt"

type ErrorWithCode interface {
	Error() string
	Code() int64
}

func Error(errcode ErrorCode) ErrorWithCode {
	return &errcode
}

func Errorf(errcode ErrorCode, format string, a ...interface{}) ErrorWithCode {
	return FromError(errcode, fmt.Errorf(format, a...))
}

func FromError(errcode ErrorCode, err error) ErrorWithCode {
	errcode.err = err
	return &errcode
}
