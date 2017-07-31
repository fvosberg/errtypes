package errtypes

import (
	"fmt"

	"github.com/pkg/errors"
)

// BadInput is used for errors, which are caused by a missing or wrong input parameter.
// Corresponding HTTP status code is 400
type BadInput interface {
	IsBadInput() bool
}

// Unauthenticated is used for errors, which are caused by missing authentication.
// Corresponding HTTP status code is 401
type Unauthenticated interface {
	IsUnauthenticated() bool
}

// Forbidden is used for errors, which are caused by unsufficient permissions
// Corresponding HTTP status code is 403
type Forbidden interface {
	IsForbidden() bool
}

// IsBadInput checks, whether this error is caused by a missing or wrong input parameter, or not
func IsBadInput(err error) bool {
	bi, ok := errors.Cause(err).(BadInput)
	return ok && bi.IsBadInput()
}

// NewBadInputError returns an error, which indicates that it's caused by a missing or wrong input parameter
func NewBadInputError(s string) error {
	return badInputError{s: s}
}

func NewBadInputErrorf(s string, i ...interface{}) error {
	return badInputError{s: fmt.Sprintf(s, i...)}
}

// badInputError is the standard implementation of the BadInput
type badInputError struct {
	s string
}

// Error returns the string representation of this error
func (e badInputError) Error() string {
	return e.s
}

// IsBadInput indicates, whether this error is caused by a missing or wrong input parameter, or not
func (e badInputError) IsBadInput() bool {
	return true
}

// IsUnauthenticated checks, whether this error is caused by a missing authentication or not
func IsUnauthenticated(err error) bool {
	bi, ok := errors.Cause(err).(Unauthenticated)
	return ok && bi.IsUnauthenticated()
}

// NewUnauthenticatedError returns an error, which indicates that it's caused by missing authentication
func NewUnauthenticatedError(s string) error {
	return unauthenticatedError{s: s}
}

// unauthenticatedError is the standard implementation of the Unauthenticated
type unauthenticatedError struct {
	s string
}

// Error returns the string representation of this error
func (e unauthenticatedError) Error() string {
	return e.s
}

// Unauthenticated indicates if this error is caused by missing authentication
func (e unauthenticatedError) IsUnauthenticated() bool {
	return true
}

// IsForbidden checks, whether this error is caused by insufficient permissions, or not
func IsForbidden(err error) bool {
	bi, ok := errors.Cause(err).(Forbidden)
	return ok && bi.IsForbidden()
}

// NewForbiddenError returns an error, which indicates that it's caused by insufficient permissions
func NewForbiddenError(s string) error {
	return forbiddenError{s: s}
}

// forbiddenError is the standard implementation of the Forbidden
type forbiddenError struct {
	s string
}

// Error returns the string representation of this error
func (e forbiddenError) Error() string {
	return e.s
}

// Forbidden indicates if this error is caused by insufficient permissions
func (e forbiddenError) IsForbidden() bool {
	return true
}
