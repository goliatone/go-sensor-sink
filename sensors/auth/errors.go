package auth

//ErrUnexpected DB unexpected error
type ErrUnexpected struct {
	debug string
}

func (err ErrUnexpected) Error() string {
	return err.debug
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

//ErrUnauthorized 401 error
type ErrUnauthorized struct {
	Message string
}

//Error string
func (err ErrUnauthorized) Error() string {
	return err.Message
}

// ErrTokenParsing ...
type ErrTokenParsing struct {
	message string
}

func (err ErrTokenParsing) Error() string {
	return err.message
}

type ErrHashPassword struct {
	password string
	message  string
}

func (err ErrHashPassword) Error() string {
	return err.message
}
