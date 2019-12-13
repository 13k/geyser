package geyser

import (
	"errors"
)

const (
	errInterfaceNotFound       = "Interface not found"
	errInvalidInterfaceName    = "Invalid interface name"
	errInterfaceMethodNotFound = "Interface method not found"
	errInvalidMethodName       = "Invalid method name"
	errInvalidMethodVersion    = "Invalid method version"
	errInvalidMethodHTTPMethod = "Invalid method HTTP method"
	errRequiredParameter       = "Missing required parameter"
)

var (
	// ErrEmptyInterfaces is returned when trying to manipulate an empty
	// SchemaInterfaces.
	ErrEmptyInterfaces = errors.New("Empty interfaces collection")
	// ErrMixedInterfaces is returned when trying to use group helpers on a
	// SchemaInterfaces collection containing mixed interface names.
	ErrMixedInterfaces = errors.New("Mixed interfaces collection")
	// ErrMixedMethods is returned when trying to use group helpers on a
	// SchemaMethods collection containing mixed method names.
	ErrMixedMethods = errors.New("Mixed methods collection")
)

// InterfaceNotFoundError is returned when tried to access an interface with
// invalid name or invalid AppID.
type InterfaceNotFoundError struct {
	Name  string
	AppID uint32
}

func (InterfaceNotFoundError) Error() string {
	return errInterfaceNotFound
}

// InvalidInterfaceNameError is returned when an interface has an invalid name.
type InvalidInterfaceNameError struct {
	Interface *SchemaInterface
	Err       error
}

func (InvalidInterfaceNameError) Error() string {
	return errInvalidInterfaceName
}

// InterfaceMethodNotFoundError is returned when tried to access an interface
// method with invalid name or invalid version.
type InterfaceMethodNotFoundError struct {
	Name    string
	Version int
}

func (InterfaceMethodNotFoundError) Error() string {
	return errInterfaceMethodNotFound
}

// InvalidMethodNameError is returned when a method has an invalid name.
type InvalidMethodNameError struct {
	Method *SchemaMethod
}

func (InvalidMethodNameError) Error() string {
	return errInvalidMethodName
}

// InvalidMethodVersionError is returned when a method has an invalid version.
type InvalidMethodVersionError struct {
	Method *SchemaMethod
}

func (InvalidMethodVersionError) Error() string {
	return errInvalidMethodVersion
}

// InvalidMethodHTTPMethodError is returned when a method has an invalid version.
type InvalidMethodHTTPMethodError struct {
	Method *SchemaMethod
}

func (InvalidMethodHTTPMethodError) Error() string {
	return errInvalidMethodHTTPMethod
}

// RequiredParameterError is returned when a required parameter is missing.
type RequiredParameterError struct {
	Param *SchemaMethodParam
}

func (RequiredParameterError) Error() string {
	return errRequiredParameter
}
