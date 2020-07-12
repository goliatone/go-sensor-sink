package device

import "fmt"

// ErrDeviceNotFound
type ErrDeviceNotFound struct {
	message string
}

func (err ErrDeviceNotFound) Error() string {
	return err.message
}

type ErrDeviceExists struct {
	message   string
	inDevice  Device
	outDevice Device
}

func (err *ErrDeviceExists) Error() string {
	err.message = "device exists"

	if err.outDevice.ID == err.inDevice.ID {
		err.message = "a device exists with the given ID"
	} else if err.outDevice.HardwareID == err.inDevice.HardwareID {
		err.message = "a device exists with the given hardware ID"
	}

	return err.message
}

type ErrUnexpected struct {
	debug string
}

func NewErrUnexpected(err error) *ErrUnexpected {
	return &ErrUnexpected{
		debug: err.Error(),
	}
}

func (err *ErrUnexpected) Error() string {
	return fmt.Sprintf("unexpected error occurred")
}

func (err *ErrUnexpected) Debug() string {
	return err.debug
}
