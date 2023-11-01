package main

import (
	"context"
	"net/http"

	"github.com/bytedance/go-tagexpr/validator"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/firewolfit/go_pkg/hzex"
	"github.com/firewolfit/go_pkg/hzex/example/handler/inner"
)

func registerRouter(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})
	testGroup := h.Group(
		"/inner",
	)
	testGroup.Any("/user", hzex.NewHandler(&inner.UserHandler{}))
}

func registerGlobalValidator() {
	validator.RegFunc("NameVd", func(args ...interface{}) error {
		return nil
	})
}
