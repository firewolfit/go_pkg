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
}

func (h *BaseHandler) MethodMiss(ctx context.Context, c *app.RequestContext, methodName string) {
}

func (h *BaseHandler) MethodExisted(ctx context.Context, c *app.RequestContext, name string) {
	c.Set("Method", name)
	c.Set("K_METHOD", name) //提供ginex框架打印metric
}

func (h *BaseHandler) Failed(ctx context.Context, c *app.RequestContext) {
	if x := recover(); x != nil {
		hlog.Errorf("%v", x)
	}
}
