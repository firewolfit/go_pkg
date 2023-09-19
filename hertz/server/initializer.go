package server

import "github.com/cloudwego/hertz/pkg/app/server"

func InitServer(headerKeys []string, queryKeys []string) server.Hertz {
	h := server.Default()
	// 设置一些东西，如：log，curl等
	h.Use(
	//....
	)
	return *h
}
