package mws

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/firewolfit/go_pkg/hzex/consts"
)

// MethodNameHandler 设置K_METHOD
func MethodNameHandler() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		methodName := GetMethodName(ctx.Query("Action"), ctx.Query("Version"))
		ctx.Set(consts.CtxHandlerMethodName, methodName)
	}
}

func GetMethodName(action, version string) string {
	if version != "" {
		version = strings.Replace(version, "-", "", -1)
	}
	return action + version
}
