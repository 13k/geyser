package geyser

import (
	"errors"
)

var (
	ErrRequestNoInterface = errors.New("geyser: no Interface set in Request")
	ErrRequestNoMethod    = errors.New("geyser: no Method set in Request")
)
