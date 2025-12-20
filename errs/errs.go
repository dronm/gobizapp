// Package errs provides error support.
package errs

import "fmt"

type PublicError interface {
	error
	Code() ErrorCode
}

type userError struct {
	msg  string    // safe message for users
	code ErrorCode // optional error code for frontend use
}

func NewPublicErrorCustom(code ErrorCode, descr string) PublicError {
	return &userError{
		code: code,
		msg:  descr,
	}
}

func NewPublicErrorWithTemplate(code ErrorCode, params ...any) PublicError {
	tmpl := ErrorDescr(code)
	descr := fmt.Sprintf(tmpl, params...)
	return &userError{
		code: code,
		msg:  descr,
	}
}

func NewPublicError(code ErrorCode) PublicError {
	descr := ErrorDescr(code)
	return &userError{
		code: code,
		msg:  descr,
	}
}

func (e *userError) Error() string {
	return e.msg
}

func (e *userError) Code() ErrorCode {
	return e.code
}

