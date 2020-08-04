package users

import (
	"net/http"
	"reflect"
	"sensors/user"
)

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ErrResponse(err error) ErrHTTP {
	var e ErrHTTP

	switch err.(type) {
	case *user.ErrRecordNotFound:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	case *user.ErrRecordExists:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Elem().Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	default:
		e = ErrHTTP{
			Error:   "Unknown",
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}
	}
	return e
}
