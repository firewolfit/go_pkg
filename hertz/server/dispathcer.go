package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/bytedance/go-tagexpr/binding"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
)

func Do(ctx context.Context, c *app.RequestContext, handler Handler) {
	var err ApiError
	var resp interface{}
	defer func() {
		DoResp(ctx, c, resp, err)
	}()

	request, _ := adaptor.GetCompatRequest(&c.Request)
	validateErr := binding.BindAndValidate(handler, request, nil)
	if validateErr != nil { // 参数绑定错误

		errDes := validateErr.Error()
		if strings.Contains(errDes, "code:") && strings.Contains(errDes, "message") {
			index := strings.Index(errDes, "code:")
			codeMessage := errDes[index:]
			code := codeMessage[6:strings.Index(codeMessage, ",")]
			message := codeMessage[strings.Index(codeMessage, "message:")+9:]
			err = NewHttpError(code, message, http.StatusBadRequest)
			return
		}

		if strings.Contains(errDes, "cause=") {
			message := errDes[strings.Index(errDes, "cause=")+6:]
			err = NewHttpError("ParamBindError", message, http.StatusBadRequest)
			return
		}
		err = NewHttpError("ParamBindError", errDes, http.StatusBadRequest)
		return
	}
	var parsableErr error
	resp, parsableErr = handler.Handle(ctx, c)
	if parsableErr != nil {
		if apiErr, ok := parsableErr.(ApiError); ok {
			err = apiErr
		} else {
			err = InternalError
		}
		return
	}
}

func DoResp(ctx context.Context, c *app.RequestContext, result interface{}, err ApiError) {
	if err != nil {
		FailedResponse(ctx, c, err)
	} else {
		SuccessResponse(ctx, c, result)
	}
}
