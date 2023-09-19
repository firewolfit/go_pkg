package server

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type IHandler interface {
	// ActionMiss 缺少Action的时候
	ActionMiss(ctx context.Context, c *app.RequestContext)
	// MethodMiss 缺少Method的时候
	MethodMiss(ctx context.Context, c *app.RequestContext)
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
	c.Set(CtxMethod, name)
}

func (h *BaseHandler) Failed(ctx context.Context, c *app.RequestContext) {
	if x := recover(); x != nil {
		hlog.Errorf("%v", x)
	}
}
