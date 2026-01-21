package convert

import "errors"

var (
	ErrUnknownUnit   = errors.New("unknown unit")
	ErrInvalidValue  = errors.New("invalid value")
	ErrNegativeValue = errors.New("negative value not allowed")
)
