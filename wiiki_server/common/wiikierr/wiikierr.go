package wiikierr

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	cause  error
	code   string // error code
	format string
	args   []interface{}
}

func (err *Error) Error() string {
	return err.cause.Error()
}

func (err *Error) Log() string {
	return fmt.Sprintf(err.format, err.args...)
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) Cause() error {
	return err.cause
}

func StackTrace(err error) {

	innerErr := err
	for {
		wiikiErr, ok := innerErr.(*Error)
		if ok {
			innerErr = wiikiErr.cause
			continue
		}

		// origin error
		fmt.Printf("%+v\n", innerErr)
		break
	}
}

func New(code string, format string, args ...interface{}) error {

	return &Error{
		cause:  errors.New(code),
		code:   code,
		format: format,
		args:   args,
	}
}

func Bind(err error, code string, format string, args ...interface{}) error {

	if !IsWiikiError(err) {
		return &Error{
			cause:  errors.WithStack(err),
			code:   code,
			format: format,
			args:   args,
		}
	}

	return &Error{
		cause:  err,
		code:   code,
		format: format,
		args:   args,
	}
}

func IsWiikiError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
