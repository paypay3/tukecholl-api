package apperrors

import (
	"fmt"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

type appError struct {
	// Error for log output
	next       error
	logMessage string
	level      level
	frame      xerrors.Frame

	// Error for gRPC api
	code    codes.Code
	message string
	details []proto.Message
}

func (e *appError) Error() string {
	if next := AsAppError(e.next); next != nil {
		return next.Error()
	}

	if e.next == nil {
		if e.logMessage != "" {
			return e.logMessage
		}

		return "no message"
	}

	return e.next.Error()
}

func (e *appError) Is(err error) bool {
	if target := AsAppError(err); target != nil {
		if e.code == target.code {
			return true
		}

		next := AsAppError(e.next)
		if next == nil {
			return false
		}

		return next.Is(err)
	}

	return false
}

func (e *appError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *appError) FormatError(p xerrors.Printer) error {
	var message string

	if e.level != "" {
		message += fmt.Sprintf("level: [%s]\n", e.level)
	}

	if e.code != codes.OK {
		message += fmt.Sprintf("code: [%s]\n", e.code.String())
	}

	if e.logMessage != "" {
		message += fmt.Sprintf("message: %s\n", e.logMessage)
	}

	if len(e.details) != 0 {
		message += fmt.Sprintf("details: %+v", e.details)
	}

	p.Print(message)
	e.frame.Format(p)
	return e.next
}

func create(msg string) *appError {
	var e appError
	e.logMessage = msg
	e.frame = xerrors.Caller(2)

	return &e
}

func New(msg string) *appError {
	return create(msg)
}

func Errorf(format string, args ...interface{}) *appError {
	return create(fmt.Sprintf(format, args...))
}

func Wrap(err error, msg ...string) *appError {
	if err == nil {
		return nil
	}

	var m string
	if len(msg) != 0 {
		m = msg[0]
	}

	e := create(m)
	e.next = err

	return e
}

func AsAppError(err error) *appError {
	if err == nil {
		return nil
	}

	var e *appError
	if xerrors.As(err, &e) {
		return e
	}

	return nil
}
