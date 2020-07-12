package sink

import "fmt"

//ErrReadingExists error for creating readings
type ErrReadingExists struct {
	message    string
	inReading  DHT22Reading
	outReading DHT22Reading
}

func (err *ErrReadingExists) Error() string {
	err.message = "reading exists"

	if err.outReading.Time == err.inReading.Time && err.outReading.Hardware == err.inReading.Hardware {
		err.message = fmt.Sprintf("duplicated reading entry")
	}

	return err.message
}

//ErrReadingNotFound error
type ErrReadingNotFound struct {
	message string
}

func (err ErrReadingNotFound) Error() string {
	return err.message
}

//ErrUnexpected error struct
type ErrUnexpected struct {
	debug string
}

//NewErrUnexpected return err unexpected
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
