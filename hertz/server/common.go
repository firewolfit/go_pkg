package server

var (
	ActionMissed  = NewHttpError("ActionMissed", "Action of the request missed", 400)
	MethodMissed  = NewHttpError("MethodMissed", "Method of the request missed", 400)
	InternalError = NewHttpError("Internal error", "Interal error occured", 500)
)
