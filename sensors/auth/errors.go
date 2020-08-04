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

//ErrRecordNotFound user not found in db
type ErrRecordNotFound struct {
	message string
}

func (err ErrRecordNotFound) Error() string {
	return err.message
}

//NewErrRecordNotFound creates a new ErrRecordNotFound error
func NewErrRecordNotFound(message string) *ErrRecordNotFound {
	return &ErrRecordNotFound{
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
