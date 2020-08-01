package authentication

import (
	"net/http"
	"reflect"
	"sensors/auth"
)

//ErrInvalidParams error
type ErrInvalidParams struct {
	message string
}

//Error print error message
func (err ErrInvalidParams) Error() string {
	return err.message
}

//ErrHTTP error
type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//ErrResponse transform Error to JSON error
func ErrResponse(err error) ErrHTTP {
	var e ErrHTTP

	switch err.(type) {
	case ErrInvalidParams:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case *auth.ErrTokenParsing:
		e = ErrHTTP{
			Error:   "ErrServerError",
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case *auth.ErrUnauthorized:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	case *auth.ErrUserNotFound:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	default:
		e = ErrHTTP{
			Error:   "Unknown",
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		}
	}

	return e
}
