package server

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/firewolfit/go_pkg/hertz/server/mws"
)

func InitServer(headerKeys []string, queryKeys []string, serviceName string, opts ...config.Option) *server.Hertz {

	responseHandler = newResponseHandler(serviceName)
	extMW := []*MiddleWare{{Func: mws.MethodNameHandler(), Index: 2}}
	return Server(extMW, opts...)
}

type MiddleWare struct {
	Func  app.HandlerFunc
	Index int
}

func Server(extMW []*MiddleWare, opts ...config.Option) *server.Hertz {

	h := server.New(opts...)

	// 插入新的middleware
	list := make([]app.HandlerFunc, 0)
	for _, newMiddleware := range extMW {
		list = InsertMiddleWare(list, newMiddleware.Index, newMiddleware.Func)
	}

	h.Use(list...)

	return h
}

func InsertMiddleWare(mws []app.HandlerFunc, index int, newMiddleware app.HandlerFunc) []app.HandlerFunc {
	if len(mws) <= index { // nil or empty slice or after last element
		return append(mws, newMiddleware)
	}
	mws = append(mws[:index+1], mws[index:]...) // Index < len(a)
	mws[index] = newMiddleware
	return mws
}
