package dsupdate

import "fmt"

// Error is Gos error extended with HTTP status and DSU Substatus.
type Error interface {
	error
	Status() int
	SubStatus() int
}

type dsuError struct {
	error
	status    int
	subStatus int
}

func (err dsuError) Status() int {
	return err.status
}

func (err dsuError) SubStatus() int {
	return err.subStatus
}

func newErrorf(status int, subStatus int, format string, a ...interface{}) Error {
	error := fmt.Errorf(format, a...)
	return dsuError{error: error, status: status, subStatus: subStatus}
}
