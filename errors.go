package geyser

const (
	errInterfaceNotFound       = "Interface not found"
	errInterfaceMethodNotFound = "Interface method not found"
	errEmptyInterfaces         = "Empty interfaces collection"
	errMixedInterfaces         = "Mixed interfaces collection"
	errMixedMethods            = "Mixed methods collection"
	errRequiredParameter       = "Missing required parameter"
)

type InterfaceNotFoundError struct {
	Name  string
	AppID uint32
}

func (InterfaceNotFoundError) Error() string {
	return errInterfaceNotFound
}

type InterfaceMethodNotFoundError struct {
	Name    string
	Version int
}

func (InterfaceMethodNotFoundError) Error() string {
	return errInterfaceMethodNotFound
}

type EmptyInterfacesError struct{}

func (EmptyInterfacesError) Error() string {
	return errEmptyInterfaces
}

type MixedInterfacesError struct{}

func (MixedInterfacesError) Error() string {
	return errMixedInterfaces
}

type MixedMethodsError struct{}

func (MixedMethodsError) Error() string {
	return errMixedMethods
}

type RequiredParameterError struct {
	Param *SchemaMethodParam
}

func (RequiredParameterError) Error() string {
	return errRequiredParameter
}
