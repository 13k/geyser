package geyser

// Schema is the root of an API schema specification.
type Schema struct {
	Host       string            `json:"host"`
	Interfaces *SchemaInterfaces `json:"interfaces"`
}

// GetInterfaces returns the underlying slice of `SchemaInterface`.
func (s *Schema) GetInterfaces() []*SchemaInterface {
	if s.Interfaces == nil {
		return nil
	}

	return s.Interfaces.Interfaces
}
