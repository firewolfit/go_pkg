package server

import (
	"context"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/firewolfit/go_pkg/hertz/server/consts"
)

type IHandler interface {
	// ActionMiss 缺少Action的时候
	ActionMiss(ctx context.Context, c *app.RequestContext)
	// MethodMiss 缺少Method的时候
	MethodMiss(ctx context.Context, c *app.RequestContext, action string)
	// MethodExisted 当存在method的时候，对methodName进行需要的操作
	MethodExisted(ctx context.Context, c *app.RequestContext, methodName string)
	// Failed 处理失败的时候
	Failed() (ctx context.Context, c *app.RequestContext)
}

type BaseHandler struct {
}

func (h *BaseHandler) ActionMiss(ctx context.Context, c *app.RequestContext) {
	FailedResponse(ctx, c, ActionMissed)
}

func (h *BaseHandler) MethodMiss(ctx context.Context, c *app.RequestContext, methodName string) {
	FailedResponse(ctx, c, MethodMissed)
}

func (h *BaseHandler) MethodExisted(ctx context.Context, c *app.RequestContext, name string) {
	c.Set(consts.CtxMethod, name)
}

func (h *BaseHandler) Failed(ctx context.Context, c *app.RequestContext) {
	if x := recover(); x != nil {
		hlog.Errorf("%v", x)
	}
}

func NewHandler(h IHandler) app.HandlerFunc {
	router := map[string]reflect.Method{}
	t := reflect.TypeOf(h)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		router[method.Name] = method
	}
	return func(ctx context.Context, c *app.RequestContext) {
		action := c.Query(consts.CtxAction)
		if action == "" {
			h.ActionMiss(ctx, c)
			return
		}
		methodName := c.GetString(consts.CtxMethod)
		if methodName == "" {
			h.MethodMiss(ctx, c, action)
		}
		h.MethodExisted(ctx, c, methodName)

		func(name string, args ...interface{}) {
			if method, ok := router[name]; ok {
				inputs := make([]reflect.Value, len(args)+1)
				inputs[0] = reflect.ValueOf(h)
				for i := range args {
					inputs[i+1] = reflect.ValueOf(args[i])
				}
				requests := method.Func.Call(inputs[0:1])
				if len(requests) != 1 {
					DoResp(ctx, c, nil, InternalError)
				}
				if request, ok := requests[0].Interface().(Handler); !ok {
					DoResp(ctx, c, nil, InternalError)
				} else {
					Do(ctx, c, request)
				}
				return
			}
			h.MethodMiss(ctx, c, name)
		}(methodName, ctx, c)
	}
}
