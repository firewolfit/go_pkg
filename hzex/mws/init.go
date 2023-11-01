package mws

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/firewolfit/go_pkg/hzex/consts"
)

func Init() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		action := ctx.Query("Action")
		version := ctx.Query("Version")
		ctx.Set(consts.CtxAction, action)
		ctx.Set(consts.CtxVersion, version)
		path := ctx.Request.Path()
		ctx.Set(consts.CtxPath, string(path))
		methodName := action + strings.Replace(version, "-", "", -1)
		ctx.Set(consts.CtxMethod, methodName)
		realIp := ctx.ClientIP()
		if xRealIp := ctx.Query("X-Real-Ip"); xRealIp != "" {
			realIp = xRealIp
		}
		ctx.Set(consts.CtxRealIp, realIp)
	}
}
