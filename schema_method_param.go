package geyser

// SchemaMethodParam holds the specification of an API interface method
// parameter.
//
// The struct should be read-only.
type SchemaMethodParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

// Validate validates the given value against the parameter specification.
//
// If the parameter is required and the value is empty, it returns an error of
// type `*RequiredParameterError`.
func (p *SchemaMethodParam) Validate(value string) error {
	if !p.Optional && value == "" {
		return &RequiredParameterError{Param: p}
	}

	return nil
}
