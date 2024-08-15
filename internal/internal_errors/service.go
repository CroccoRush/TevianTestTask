package internal_errors

import "errors"

var (
	ErrInvalidJson    = errors.New("invalid json body")
	ErrLocked         = errors.New("locked")
	ErrAlreadyProcess = errors.New("already process")
)
