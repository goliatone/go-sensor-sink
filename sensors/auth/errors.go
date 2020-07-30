package auth

//ErrUnexpected DB unexpected error
type ErrUnexpected struct {
	debug string
}

//NewErrUnexpected creates a new ErrUnexpected DB unexpected error
func NewErrUnexpected(err error) *ErrUnexpected {
	return &ErrUnexpected{
		debug: err.Error(),
	}
}

//ErrUserNotFound user not found in db
type ErrUserNotFound struct {
	message string
}

func (err ErrUserNotFound) Error() string {
	return err.message
}

//NewErrUserNotFound creates a new ErrUserNotFound error
func NewErrUserNotFound(message string) *ErrUserNotFound {
	return &ErrUserNotFound{
		message: message,
	}
}
