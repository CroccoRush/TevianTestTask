package internal_errors

import "errors"

var (
	ErrInvalidParams = errors.New("invalid params")
	ErrNotFound      = errors.New("not found")
	ErrInternal      = errors.New("internal error")
	ErrUnexpected    = errors.New("unexpected error")
)
