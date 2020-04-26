package schema

// Schema is the root of an API schema specification.
type Schema struct {
	Host       string     `json:"host"`
	Interfaces Interfaces `json:"interfaces"`
}

// Validate checks if all contained interfaces are valid.
//
// Returns errors described in `Interfaces.Validate`
func (s *Schema) Validate() error {
	return s.Interfaces.Validate()
}
