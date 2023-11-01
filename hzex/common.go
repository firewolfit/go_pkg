package hzex

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type Handler interface {
	Handle(ctx context.Context, c *app.RequestContext) (interface{}, error)
}

var (
	ActionMissed   = NewHttpError("ActionMissed", "Action of the request missed", http.StatusBadRequest)
	MethodMissed   = NewHttpError("MethodMissed", "Method of the request missed", http.StatusBadRequest)
	InternalError  = NewHttpError("Internal error", "Interal error occured", http.StatusInternalServerError)
	UndefinedError = NewHttpError("UndefinedError", "Undefined error orrcured", http.StatusInternalServerError)
)
