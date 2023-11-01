package hzex

import "fmt"

type ApiError interface {
	Error() string
	HttpStatus() int
	BizErrCode() string
	BizErrCodeInt() int
	GetMessage() string
	GetArgs() []interface{}
	WithArgs(args ...interface{}) ApiError
}

type HttpError struct {
	HttpCode int
	Code     string
	CodeInt  int
	Message  string
	Args     []interface{}
}

func (he *HttpError) Error() string {
	msg := fmt.Sprintf(he.Message, he.Args...)
	return fmt.Sprintf("code: %s, message: %s", he.Code, msg)
}
func (he *HttpError) HttpStatus() int {
	return he.HttpCode
}

func (he *HttpError) BizErrCode() string {
	return he.Code
}

func (he *HttpError) BizErrCodeInt() int {
	return he.CodeInt
}

func (he *HttpError) GetMessage() string {
	return he.Message
}
func (he *HttpError) GetArgs() []interface{} {
	return he.Args
}
func (he *HttpError) WithArgs(args ...interface{}) ApiError {
	e2 := *he
	e2.Args = append(e2.Args, args...)
	return &e2
}

func NewHttpError(code, msg string, httpCode int) ApiError {
	return &HttpError{Code: code, Message: msg, HttpCode: httpCode}
}
