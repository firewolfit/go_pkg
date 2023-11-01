package inner

import (
	"github.com/firewolfit/go_pkg/hzex"
	"github.com/firewolfit/go_pkg/hzex/example/controller/inner/user"
)

type UserHandler struct {
	hzex.BaseHandler
}

func (*UserHandler) GetUser20180101() hzex.Handler {
	return &user.GetUser{}
}
