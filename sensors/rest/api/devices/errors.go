package devices

import (
	"net/http"
	"reflect"
	"sensors/device"
)

//ErrHTTP response rror
type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//ErrResponse transform an Error into a JSON error response
func ErrResponse(err error) ErrHTTP {
	var e ErrHTTP

	switch err.(type) {
	case *device.ErrDeviceNotFound:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	case *device.ErrDeviceExists:
		e = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
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
