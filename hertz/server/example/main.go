package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/firewolfit/go_pkg/hertz/server"
)

func main() {
	hertzServer := server.InitServer(nil, nil, "test")
	hertzServer.Group(
		"/test", func(c context.Context, ctx *app.RequestContext) {
			fmt.Println("xxxx")
		},
	)
	hertzServer.Spin()
}
