package user

import "fmt"

// ErrRecordNotFound
type ErrRecordNotFound struct {
	message string
}

func (err ErrRecordNotFound) Error() string {
	return err.message
}

//ErrRecordExists when we insert duplicated record
type ErrRecordExists struct {
	message string
	inUser  User
	outUser User
}

func (err *ErrRecordExists) Error() string {
	err.message = "User exists"

	if err.outUser.ID == err.inUser.ID {
		err.message = "A User exists with the given ID"
	}
	return err.message
}

//ErrUnexpected unkown error
type ErrUnexpected struct {
	debug string
}

func NewErrUnexpected(err error) *ErrUnexpected {
	return &ErrUnexpected{
		debug: err.Error(),
	}
}

func (err *ErrUnexpected) Error() string {
	return fmt.Sprintf("Unexpected error ocurred")
}

func (err *ErrUnexpected) Debug() string {
	return err.debug
}
