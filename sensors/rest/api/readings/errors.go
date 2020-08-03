package readings

import (
	"net/http"
	"reflect"
	"sensors/sink"
)

//ErrHTTP response error
type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//ErrResponse transform an Error into a JSON error response
func ErrResponse(err error) ErrHTTP {
	var e ErrHTTP

	switch err.(type) {
	case *sink.ErrReadingExists:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Elem().Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case *sink.ErrReadingNotFound:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Elem().Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case *sink.ErrUnexpected:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Elem().Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	default:
		e = ErrHTTP{
			Error:   "Unknown Error",
			Message: "Internanal server error",
			Status:  http.StatusInternalServerError,
		}
	}

	return e
}
