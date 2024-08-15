package internal_errors

import "errors"

var (
	ErrUnexpectedDB = errors.New("unexpected database error")
	ErrDuplicateKey = errors.New("duplicate key")
)
