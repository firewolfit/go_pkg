package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetUser struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func (user *GetUser) Handle(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	return user, nil
}
