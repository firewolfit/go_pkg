package server

import "net/http"

var (
	ActionMissed   = NewHttpError("ActionMissed", "Action of the request missed", http.StatusBadRequest)
	MethodMissed   = NewHttpError("MethodMissed", "Method of the request missed", http.StatusBadRequest)
	InternalError  = NewHttpError("Internal error", "Interal error occured", http.StatusInternalServerError)
	UndefinedError = NewHttpError("UndefinedError", "Undefined error orrcured", http.StatusInternalServerError)
)
