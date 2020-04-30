package geyser

import (
	"errors"
)

var (
	// ErrRequestNoInterface is returned when trying to execute a `Request` without an `Interface` set.
	ErrRequestNoInterface = errors.New("geyser: no Interface set in Request")
	// ErrRequestNoMethod is returned when trying to execute a `Request` without an `Method` set.
	ErrRequestNoMethod = errors.New("geyser: no Method set in Request")
)
