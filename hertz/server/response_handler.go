package server

import (
	"context"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
)

type ResponseHandler struct {
	defaultError ApiError
}

func NewResponseHandler(serviceName string) *ResponseHandler {
	handler := &ResponseHandler{
		defaultError: UndefinedError,
	}
	return handler
}

func (*ResponseHandler) getRequestId(ctx context.Context, c *app.RequestContext) string {
	return c.GetString(CtxRequestId)
}

func (*ResponseHandler) getAction(ctx context.Context, c *app.RequestContext) string {
	return c.GetString(CtxAction)
}

func (*ResponseHandler) getVersion(ctx context.Context, c *app.RequestContext) string {
	return c.GetString(CtxVersion)
}

// doResponse 返回
func (*ResponseHandler) doResponse(ctx context.Context, c *app.RequestContext, httpCode, bizCodeInt int, bizErrCode, bizErrMsg string, result interface{}) {
	if result != nil {
		if v := reflect.ValueOf(result); v.Kind() == reflect.Ptr && v.IsNil() {
			result = nil
		}
	}
	res := &Response{
		Meta:   ResponseMeta{},
		Result: result,
	}
	if bizErrCode != "" {
		res.Meta.Error = &HttpError{
			CodeInt: bizCodeInt,
			Code:    bizErrCode,
			Message: bizErrMsg,
		}
		c.Abort()
		c.Set(CtxErrorFlag, 1)
	}
	c.JSON(httpCode, res)
}
