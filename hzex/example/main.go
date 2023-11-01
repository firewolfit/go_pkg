package main

import (
	"github.com/firewolfit/go_pkg/hzex"
)

func main() {
	h := hzex.InitServer()
	registerRouter(h)
	h.Spin()
}
