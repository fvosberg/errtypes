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

// NotFound is used, when a requested recource doesn't exist
// Corresponding HTTP status code is 404
type NotFound interface {
	IsNotFound() bool
}

// IsBadInput checks, whether this error is caused by a missing or wrong input parameter, or not
func IsBadInput(err error) bool {
	bi, ok := errors.Cause(err).(BadInput)
	return ok && bi.IsBadInput()
}

// NewBadInputError returns an error, which indicates that it's caused by a missing or wrong input parameter
func NewBadInput(s string) error {
	return badInputError{s: s}
}

func NewBadInputf(s string, i ...interface{}) error {
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

// NewUnauthenticated returns an error, which indicates that it's caused by missing authentication
func NewUnauthenticated(s string) error {
	return unauthenticatedError{s: s}
}

// NewUnauthenticatedf returns an error, which indicates that it's caused by missing authentication
// it accepts a format string and a variadic argument for it
func NewUnauthenticatedf(s string, args ...interface{}) error {
	return unauthenticatedError{s: fmt.Sprintf(s, args...)}
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

// NewForbidden returns an error, which indicates that it's caused by insufficient permissions
func NewForbidden(s string) error {
	return forbiddenError{s: s}
}

// NewForbiddenf returns an error, which indicates that it's caused by insufficient permissions
func NewForbiddenf(s string, i ...interface{}) error {
	return forbiddenError{s: fmt.Sprintf(s, i...)}
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

// IsNotFound checks, whether this error is caused by a missing resource
func IsNotFound(err error) bool {
	bi, ok := errors.Cause(err).(NotFound)
	return ok && bi.IsNotFound()
}

// NewNotFound returns an error, which indicates that it's caused by a missing resource
func NewNotFound(s string) error {
	return notFoundError{s: s}
}

// NewNotFoundf returns an error, which indicates that it's caused by a missing resource - supports sprintf
func NewNotFoundf(s string, i ...interface{}) error {
	return notFoundError{s: fmt.Sprintf(s, i...)}
}

// notFoundError is the standard implementation of the NotFound
type notFoundError struct {
	s string
}

// Error returns the string representation of this error
func (e notFoundError) Error() string {
	return e.s
}

// NotFound indicates if this error is caused by a missing resource
func (e notFoundError) IsNotFound() bool {
	return true
}
