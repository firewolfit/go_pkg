package server

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

var responseHandler *ResponseHandler

const (
	CtxErrorFlag = "__has_error__"
)

// ResponseMeta 返回元数据
type ResponseMeta struct {
	RequestId string   `json:",omitempty"`
	Action    string   `json:",omitempty"`
	Version   string   `json:",omitempty"`
	Error     ApiError `json:",omitempty"`
}

// Response 返回值
type Response struct {
	Meta   ResponseMeta `json:",omitempty"`
	Result interface{}  `json:",omitempty"`
}

// SuccessResponse 处理成功返回
func SuccessResponse(ctx context.Context, c *app.RequestContext, result interface{}) {
	responseHandler.doResponse(ctx, c, http.StatusOK, 0, "", "", result)
}

// FailedResponse 错误返回
func FailedResponse(ctx context.Context, c *app.RequestContext, err ApiError) {
	if err == nil {
		err = responseHandler.defaultError
	}
	responseHandler.doResponse(ctx, c, err.HttpStatus(), err.BizErrCodeInt(), err.BizErrCode(), err.GetMessage(), nil)
}
