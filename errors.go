package geyser

const (
	errInterfaceNotFound       = "Interface not found"
	errInterfaceMethodNotFound = "Interface method not found"
	errEmptyInterfaces         = "Empty interfaces collection"
	errMixedInterfaces         = "Mixed interfaces collection"
	errMixedMethods            = "Mixed methods collection"
	errRequiredParameter       = "Missing required parameter"
)

// InterfaceNotFoundError is returned when tried to access an interface with invalid name or invalid
// AppID.
type InterfaceNotFoundError struct {
	Name  string
	AppID uint32
}

func (InterfaceNotFoundError) Error() string {
	return errInterfaceNotFound
}

// InterfaceMethodNotFoundError is returned when tried to access an interface method with invalid
// name or invalid version.
type InterfaceMethodNotFoundError struct {
	Name    string
	Version int
}

func (InterfaceMethodNotFoundError) Error() string {
	return errInterfaceMethodNotFound
}

// EmptyInterfacesError is returned when trying to manipulate an empty SchemaInterfaces.
type EmptyInterfacesError struct{}

func (EmptyInterfacesError) Error() string {
	return errEmptyInterfaces
}

// MixedInterfacesError is returned when trying to use group helpers on a SchemaInterfaces
// collection containing mixed interface names.
type MixedInterfacesError struct{}

func (MixedInterfacesError) Error() string {
	return errMixedInterfaces
}

// MixedMethodsError is returned when trying to use group helpers on a SchemaMethods collection
// containing mixed method names.
type MixedMethodsError struct{}

func (MixedMethodsError) Error() string {
	return errMixedMethods
}

// RequiredParameterError is returned when a required parameter is missing.
type RequiredParameterError struct {
	Param *SchemaMethodParam
}

func (RequiredParameterError) Error() string {
	return errRequiredParameter
}
